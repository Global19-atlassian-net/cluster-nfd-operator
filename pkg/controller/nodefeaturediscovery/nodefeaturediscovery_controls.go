package nodefeaturediscovery

import (
	"context"

	secv1 "github.com/openshift/api/security/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	nfdconfig "github.com/openshift/cluster-nfd-operator/pkg/config"
)

type controlFunc []func(n NFD) (ResourceStatus, error)

type ResourceStatus int

const (
	Ready    ResourceStatus = 0
	NotReady ResourceStatus = 1
)

func (s ResourceStatus) String() string {
	names := [...]string{
		"Ready",
		"NotReady"}

	if s < Ready || s > NotReady {
		return "Unkown Resources Status"
	}
	return names[s]
}

func Namespace(n NFD) (ResourceStatus, error) {

	log.Info("Namespace0")

	state := n.idx
	obj := n.resources[state].Namespace

	log.Info("Namespace1")

	if len(n.ins.Spec.OperandNamespace) != 0 {
		obj.SetName(n.ins.Spec.OperandNamespace)
	}
	log.Info("Namespace2")

	found := &corev1.Namespace{}
	logger := log.WithValues("Namespace", obj.Name, "Namespace", "Cluster")

	log.Info("Namespace3")

	if err := controllerutil.SetControllerReference(n.ins, &obj, n.rec.scheme); err != nil {
		return NotReady, err
	}
	log.Info("Namespace4")

	logger.Info("Looking for")
	err := n.rec.client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating ")
		err = n.rec.client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, skpping update")

	return Ready, nil
}

func ServiceAccount(n NFD) (ResourceStatus, error) {

	state := n.idx
	obj := n.resources[state].ServiceAccount

	if len(n.ins.Spec.OperandNamespace) != 0 {
		obj.SetNamespace(n.ins.Spec.OperandNamespace)
	}

	found := &corev1.ServiceAccount{}
	logger := log.WithValues("ServiceAccount", obj.Name, "Namespace", obj.Namespace)

	if err := controllerutil.SetControllerReference(n.ins, &obj, n.rec.scheme); err != nil {
		return NotReady, err
	}

	logger.Info("Looking for")
	err := n.rec.client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating ")
		err = n.rec.client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, skpping update")

	return Ready, nil
}
func ClusterRole(n NFD) (ResourceStatus, error) {

	state := n.idx
	obj := n.resources[state].ClusterRole

	if len(n.ins.Spec.OperandNamespace) != 0 {
		obj.SetNamespace(n.ins.Spec.OperandNamespace)
	}

	found := &rbacv1.ClusterRole{}
	logger := log.WithValues("ClusterRole", obj.Name, "Namespace", obj.Namespace)

	if err := controllerutil.SetControllerReference(n.ins, &obj, n.rec.scheme); err != nil {
		return NotReady, err
	}

	logger.Info("Looking for")
	err := n.rec.client.Get(context.TODO(), types.NamespacedName{Namespace: "", Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = n.rec.client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")
	err = n.rec.client.Update(context.TODO(), &obj)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}

func ClusterRoleBinding(n NFD) (ResourceStatus, error) {

	state := n.idx
	obj := n.resources[state].ClusterRoleBinding

	if len(n.ins.Spec.OperandNamespace) != 0 {
		obj.SetNamespace(n.ins.Spec.OperandNamespace)
	}

	found := &rbacv1.ClusterRoleBinding{}
	logger := log.WithValues("ClusterRoleBinding", obj.Name, "Namespace", obj.Namespace)

	if err := controllerutil.SetControllerReference(n.ins, &obj, n.rec.scheme); err != nil {
		return NotReady, err
	}

	logger.Info("Looking for")
	err := n.rec.client.Get(context.TODO(), types.NamespacedName{Namespace: "", Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = n.rec.client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")
	err = n.rec.client.Update(context.TODO(), &obj)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}
func Role(n NFD) (ResourceStatus, error) {

	state := n.idx
	obj := n.resources[state].Role

	if len(n.ins.Spec.OperandNamespace) != 0 {
		obj.SetNamespace(n.ins.Spec.OperandNamespace)
	}

	found := &rbacv1.Role{}
	logger := log.WithValues("Role", obj.Name, "Namespace", obj.Namespace)

	if err := controllerutil.SetControllerReference(n.ins, &obj, n.rec.scheme); err != nil {
		return NotReady, err
	}

	logger.Info("Looking for")
	err := n.rec.client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = n.rec.client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")
	err = n.rec.client.Update(context.TODO(), &obj)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}

func RoleBinding(n NFD) (ResourceStatus, error) {

	state := n.idx
	obj := n.resources[state].RoleBinding

	if len(n.ins.Spec.OperandNamespace) != 0 {
		obj.SetNamespace(n.ins.Spec.OperandNamespace)
	}

	found := &rbacv1.RoleBinding{}
	logger := log.WithValues("RoleBinding", obj.Name, "Namespace", obj.Namespace)

	if err := controllerutil.SetControllerReference(n.ins, &obj, n.rec.scheme); err != nil {
		return NotReady, err
	}

	logger.Info("Looking for")
	err := n.rec.client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = n.rec.client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")
	err = n.rec.client.Update(context.TODO(), &obj)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}

func ConfigMap(n NFD) (ResourceStatus, error) {

	state := n.idx
	obj := n.resources[state].ConfigMap

	if len(n.ins.Spec.OperandNamespace) != 0 {
		obj.SetNamespace(n.ins.Spec.OperandNamespace)
	}

	found := &corev1.ConfigMap{}
	logger := log.WithValues("ConfigMap", obj.Name, "Namespace", obj.Namespace)

	if err := controllerutil.SetControllerReference(n.ins, &obj, n.rec.scheme); err != nil {
		return NotReady, err
	}

	logger.Info("Looking for")
	err := n.rec.client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = n.rec.client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")
	err = n.rec.client.Update(context.TODO(), &obj)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}

func DaemonSet(n NFD) (ResourceStatus, error) {

	state := n.idx
	obj := n.resources[state].DaemonSet

	if len(n.ins.Spec.OperandNamespace) != 0 {
		obj.SetNamespace(n.ins.Spec.OperandNamespace)
	}

	obj.Spec.Template.Spec.Containers[0].Image = nfdconfig.NodeFeatureDiscoveryImage()

	if len(n.ins.Spec.OperandImage) != 0 {
		obj.Spec.Template.Spec.Containers[0].Image = n.ins.Spec.OperandImage
	}

	found := &appsv1.DaemonSet{}
	logger := log.WithValues("DaemonSet", obj.Name, "Namespace", obj.Namespace)

	if err := controllerutil.SetControllerReference(n.ins, &obj, n.rec.scheme); err != nil {
		return NotReady, err
	}

	logger.Info("Looking for")
	err := n.rec.client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = n.rec.client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")
	err = n.rec.client.Update(context.TODO(), &obj)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}

func Service(n NFD) (ResourceStatus, error) {

	state := n.idx
	obj := n.resources[state].Service

	if len(n.ins.Spec.OperandNamespace) != 0 {
		obj.SetNamespace(n.ins.Spec.OperandNamespace)
	}

	found := &corev1.Service{}
	logger := log.WithValues("Service", obj.Name, "Namespace", obj.Namespace)

	if err := controllerutil.SetControllerReference(n.ins, &obj, n.rec.scheme); err != nil {
		return NotReady, err
	}

	logger.Info("Looking for")
	err := n.rec.client.Get(context.TODO(), types.NamespacedName{Namespace: obj.Namespace, Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = n.rec.client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create")
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")

	required := obj.DeepCopy()
	required.ResourceVersion = found.ResourceVersion
	required.Spec.ClusterIP = found.Spec.ClusterIP

	err = n.rec.client.Update(context.TODO(), required)

	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}

func SecurityContextConstraints(n NFD) (ResourceStatus, error) {

	state := n.idx
	obj := n.resources[state].SecurityContextConstraints

	if len(n.ins.Spec.OperandNamespace) != 0 {
		obj.SetNamespace(n.ins.Spec.OperandNamespace)
	}

	found := &secv1.SecurityContextConstraints{}
	logger := log.WithValues("SecurityContextConstraints", obj.Name, "Namespace", "default")

	if err := controllerutil.SetControllerReference(n.ins, &obj, n.rec.scheme); err != nil {
		return NotReady, err
	}

	logger.Info("Looking for")
	err := n.rec.client.Get(context.TODO(), types.NamespacedName{Namespace: "", Name: obj.Name}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Not found, creating")
		err = n.rec.client.Create(context.TODO(), &obj)
		if err != nil {
			logger.Info("Couldn't create", "Error", err)
			return NotReady, err
		}
		return Ready, nil
	} else if err != nil {
		return NotReady, err
	}

	logger.Info("Found, updating")

	required := obj.DeepCopy()
	required.ResourceVersion = found.ResourceVersion

	err = n.rec.client.Update(context.TODO(), required)
	if err != nil {
		return NotReady, err
	}

	return Ready, nil
}
