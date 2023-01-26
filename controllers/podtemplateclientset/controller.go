package podtemplateclientset

import (
	"context"

	"github.com/INFURA/infrakit/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"

	infrakitv1alpha1 "github.com/INFURA/infrakit/api/v1alpha1"
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
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MainReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrakitv1alpha1.PodTemplateClientSet{}).
		Complete(r)
}
