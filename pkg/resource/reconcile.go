package resource

import (
	"context"
	"fmt"
	"hash/fnv"
	"strings"
	"time"

	"github.com/INFURA/infrakit/pkg/util/hash"
	"github.com/avast/retry-go"
	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// HashLabel is a common label for each resource.
	HashLabel = "common.infura.io/hash"
	// DefaultRetryAttempts is the number of retry when deleting resource.
	DefaultRetryAttempts = 10
	// DefaultRetryDelay is the duration between each retry when deleting resource.
	DefaultRetryDelay = 1 * time.Second
)

type DeleteStatefulSetError struct {
	msg string
	err error
}

func (d *DeleteStatefulSetError) Error() string {
	return fmt.Sprintf("%s: %v", d.msg, d.err)
}

func NewDeleteStatefulSetError(msg string, err error) error {
	return &DeleteStatefulSetError{msg: msg, err: err}
}

type ReconciliableObject interface {
	metav1.Object
	runtime.Object
}

// SetOwnerAndHash sets an owner reference and hash annotation on the object.
func SetOwnerAndHash(owner ReconciliableObject, object metav1.Object, scheme *runtime.Scheme) error {
	err := ctrl.SetControllerReference(owner, object, scheme)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	objectHash := makeHash(object)

	objectAnnotations := object.GetAnnotations()
	if objectAnnotations == nil {
		objectAnnotations = make(map[string]string)
	}

	objectAnnotations[HashLabel] = objectHash
	object.SetAnnotations(objectAnnotations)

	return nil
}

// ReconcileObject reconciles an object, either creating or updating if needed.
func ReconcileObject(
	ctx context.Context,
	log logr.Logger,
	client client.Client,
	object ReconciliableObject,
	foundObject ReconciliableObject,
) error {
	name := object.GetName()
	namespace := object.GetNamespace()
	namespacedName := types.NamespacedName{Name: name, Namespace: namespace}

	err := client.Get(ctx, namespacedName, foundObject)

	switch {
	case err != nil && errors.IsNotFound(err):
		{
			log.Info("Creating a new object",
				"Type", fmt.Sprintf("%T", object),
				"Namespace", object.GetNamespace(),
				"Name", object.GetName(),
			)
			err = client.Create(ctx, object)
			if err != nil {
				return fmt.Errorf("error creating object [%s/%s]: %w", namespace, name, err)
			}
		}
	case err != nil:
		{
			return fmt.Errorf("error getting object [%s/%s]: %w", namespace, name, err)
		}
	default:
		{
			objectHash := object.GetAnnotations()[HashLabel]
			foundObjectHash := foundObject.GetAnnotations()[HashLabel]

			if foundObjectHash != objectHash {
				object.SetResourceVersion(foundObject.GetResourceVersion())
				log.Info("Updating an object",
					// "Kind", object.GetObjectKind().GroupVersionKind().GroupKind().String(),
					"Namespace", object.GetNamespace(), "Name", object.GetName(),
					"Hash", objectHash, "OldHash", foundObjectHash)
				err = client.Update(ctx, object)

				if err != nil {
					return fmt.Errorf("error updating object [%s/%s]: %w", namespace, name, err)
				}
			}
		}
	}

	return nil
}

func ReconcileStatefulSet(
	ctx context.Context,
	log logr.Logger,
	clientIn client.Client,
	owner ReconciliableObject,
	object *appsv1.StatefulSet,
	scheme *runtime.Scheme,
) error {
	existingStatefulSet := &appsv1.StatefulSet{}
	namespacedName := types.NamespacedName{Name: object.Name, Namespace: object.Namespace}

	err := clientIn.Get(ctx, namespacedName, existingStatefulSet)
	if client.IgnoreNotFound(err) != nil {
		return fmt.Errorf("%w", err)
	}

	if !errors.IsNotFound(err) { //nolint: nestif
		comparer := NewStatefulSetComparer(*existingStatefulSet)

		result, err := comparer.Compare(*object)
		if err != nil {
			return fmt.Errorf("error comparing statefulsets: %w", err)
		}

		if !result.Match() {
			msg := "StatefulSets differ"

			if result.Replace() {
				msg += ", need to replace"
			} else {
				msg += ", updating in place"
			}

			log.V(0).Info(msg, "diff", strings.Join(result.Reasons(), "; "))
		}

		if result.Replace() {
			// Delete the current statefulset without deleting the pods
			err := DeleteStatefulSet(ctx, clientIn, existingStatefulSet, metav1.DeletePropagationOrphan)
			if err != nil {
				return err
			}
		}
	}

	err = SetOwnerAndHash(owner, object, scheme)
	if err != nil {
		return fmt.Errorf("failed to reconcile statefulset: %w", err)
	}

	foundObject := &appsv1.StatefulSet{}

	return ReconcileObject(ctx, log, clientIn, object, foundObject)
}

// DeleteStatefulSet deletes the provided statefulset and waits until the object is cleared.
func DeleteStatefulSet(
	ctx context.Context,
	clientIn client.Client,
	statefulSet *appsv1.StatefulSet,
	propagationPolicy metav1.DeletionPropagation,
) error {
	if statefulSet == nil {
		return NewDeleteStatefulSetError("there is no statefulset in the cluster", nil)
	}

	options := &client.DeleteOptions{PropagationPolicy: &propagationPolicy}

	err := clientIn.Delete(ctx, statefulSet, options)
	if err != nil {
		return NewDeleteStatefulSetError("could not delete statefulset: %v", err)
	}

	objKey := client.ObjectKeyFromObject(statefulSet)

	// Ensure deleted
	err = retry.Do(
		func() error {
			err := clientIn.Get(ctx, objKey, &appsv1.StatefulSet{})
			if errors.IsNotFound(err) {
				return nil
			}

			// Found it fine, this is an error for us
			if err == nil {
				return NewDeleteStatefulSetError("statefulset still deleting: %v", err)
			}

			return fmt.Errorf("%w", err)
		},
		retry.Attempts(DefaultRetryAttempts),
		retry.Delay(DefaultRetryDelay),
		retry.DelayType(retry.FixedDelay),
	)
	if err != nil {
		return NewDeleteStatefulSetError("could not delete statefulset: %v", err)
	}

	return nil
}

func makeHash(object interface{}) string {
	hf := fnv.New32()
	hash.DeepHashObject(hf, object)

	return fmt.Sprint(hf.Sum32())
}
