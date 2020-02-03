package nodefeaturediscovery

import (
	"context"

	secv1 "github.com/openshift/api/security/v1"
	nfdv1alpha1 "github.com/openshift/cluster-nfd-operator/pkg/apis/nfd/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_nodefeaturediscovery")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new NodeFeatureDiscovery Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))

}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileNodeFeatureDiscovery{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("nodefeaturediscovery-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource NodeFeatureDiscovery
	if err = c.Watch(&source.Kind{Type: &nfdv1alpha1.NodeFeatureDiscovery{}}, &handler.EnqueueRequestForObject{}); err != nil {
		return err
	}

	// Watch for changes to secondary resource Pods and requeue the owner NodeFeatureDiscovery
	cache := &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &nfdv1alpha1.NodeFeatureDiscovery{},
	}

	if err = c.Watch(&source.Kind{Type: &appsv1.DaemonSet{}}, cache); err != nil {
		return err
	}

	if err = c.Watch(&source.Kind{Type: &corev1.Service{}}, cache); err != nil {
		return err
	}

	if err = c.Watch(&source.Kind{Type: &corev1.ServiceAccount{}}, cache); err != nil {
		return err
	}

	if err = c.Watch(&source.Kind{Type: &secv1.SecurityContextConstraints{}}, cache); err != nil {
		return err
	}

	if err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, cache); err != nil {
		return err
	}

	if err = c.Watch(&source.Kind{Type: &corev1.ConfigMap{}}, cache); err != nil {
		return err
	}

	if err = c.Watch(&source.Kind{Type: &rbacv1.Role{}}, cache); err != nil {
		return err
	}

	if err = c.Watch(&source.Kind{Type: &rbacv1.RoleBinding{}}, cache); err != nil {
		return err
	}
	if err = c.Watch(&source.Kind{Type: &rbacv1.ClusterRole{}}, cache); err != nil {
		return err
	}

	if err = c.Watch(&source.Kind{Type: &rbacv1.ClusterRoleBinding{}}, cache); err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileNodeFeatureDiscovery{}
var nfd NFD

// ReconcileNodeFeatureDiscovery reconciles a NodeFeatureDiscovery object
type ReconcileNodeFeatureDiscovery struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a NodeFeatureDiscovery object and makes changes based on the state read
// and what is in the NodeFeatureDiscovery.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileNodeFeatureDiscovery) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling NodeFeatureDiscovery")

	reqLogger.Info("DEBUG TEST TEST TEST")

	// Fetch the NodeFeatureDiscovery instance
	instance := &nfdv1alpha1.NodeFeatureDiscovery{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	nfd.init(r, instance)

	for {
		err := nfd.step()
		if err != nil {
			return reconcile.Result{}, err
		}
		if nfd.last() {
			break
		}
	}

	reqLogger.Info("DEBUG TEST TEST TEST22222")

	return reconcile.Result{Requeue: false}, nil
}
