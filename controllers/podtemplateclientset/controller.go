package podtemplateclientset

import (
	"context"
	"fmt"

	infrakitv1alpha1 "github.com/INFURA/infrakit/api/v1alpha1"
	"github.com/INFURA/infrakit/controllers"
	"github.com/INFURA/infrakit/pkg/resource"
	"github.com/INFURA/infrakit/pkg/util/label"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// MainReconciler reconciles a PodTemplateClientSet object
type MainReconciler struct {
	controllers.Reconciler
}

//+kubebuilder:rbac:groups=infrakit.infura.io,resources=podtemplateclientsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=infrakit.infura.io,resources=podtemplateclientsets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=infrakit.infura.io,resources=podtemplateclientsets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PodTemplateClientSet object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile
func (r *MainReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var ptcs infrakitv1alpha1.PodTemplateClientSet

	err := r.Client.Get(ctx, req.NamespacedName, &ptcs)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err) //nolint: wrapcheck
	}

	// TODO: Apply the default settings
	// ptcs.Default()

	// Union labels on the PTCS
	operatorManagedLabels := SubresourceLabels(&ptcs)
	ptcs.Spec.Template.ObjectMeta.Labels = label.Union(operatorManagedLabels, ptcs.Spec.Template.ObjectMeta.Labels)

	sts := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ptcs.Name,
			Namespace: ptcs.Namespace,
		},
		Spec: appsv1.StatefulSetSpec{
			Selector:            metav1.SetAsLabelSelector(ptcs.Spec.Template.ObjectMeta.Labels),
			PodManagementPolicy: appsv1.ParallelPodManagement,
			Replicas:            ptcs.Spec.Replicas,
			Template:            ptcs.Spec.Template,
		},
	}

	err = resource.ReconcileStatefulSet(ctx, log, r.Client, &ptcs, sts, r.Scheme)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("error reconciling statefulset: %w", err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MainReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrakitv1alpha1.PodTemplateClientSet{}).
		Complete(r)
}
