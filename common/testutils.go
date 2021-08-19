package common

import (
	infrastructurev1alpha4 "github.com/vmware-tanzu/cluster-api-provider-byoh/apis/infrastructure/v1alpha4"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

func NewByoMachine(byoMachineName string, byoMachineNamespace string, clusterName string, machine *clusterv1.Machine) *infrastructurev1alpha4.ByoMachine {

	byoMachine := &infrastructurev1alpha4.ByoMachine{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ByoMachine",
			APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha4",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      byoMachineName,
			Namespace: byoMachineNamespace,
		},
		Spec: infrastructurev1alpha4.ByoMachineSpec{},
	}

	if machine != nil {
		byoMachine.ObjectMeta.OwnerReferences = []metav1.OwnerReference{
			{
				Kind:       "Machine",
				Name:       machine.Name,
				APIVersion: "cluster.x-k8s.io/v1",
				UID:        machine.UID,
			},
		}
	}

	if len(clusterName) > 0 {
		byoMachine.ObjectMeta.Labels = map[string]string{
			clusterv1.ClusterLabelName: clusterName,
		}
	}
	return byoMachine
}

func NewMachine(machineName string, namespace string, clusterName string) *clusterv1.Machine {
	machine := &clusterv1.Machine{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Machine",
			APIVersion: "cluster.x-k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      machineName,
			Namespace: namespace,
		},
		Spec: clusterv1.MachineSpec{
			ClusterName: clusterName,
		},
	}
	return machine
}

func NewByoHost(byoHostName string, byoHostNamespace string, byoMachine *infrastructurev1alpha4.ByoMachine) *infrastructurev1alpha4.ByoHost {
	byoHost := &infrastructurev1alpha4.ByoHost{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ByoHost",
			APIVersion: "infrastructure.cluster.x-k8s.io/v1alpha4",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      byoHostName,
			Namespace: byoHostNamespace,
		},
		Spec: infrastructurev1alpha4.ByoHostSpec{},
	}

	if byoMachine != nil {
		byoHost.Status.MachineRef = &corev1.ObjectReference{
			Kind:       "ByoMachine",
			Namespace:  byoMachine.Namespace,
			Name:       byoMachine.Name,
			UID:        byoMachine.UID,
			APIVersion: byoHost.APIVersion,
		}
	}
	return byoHost
}

func NewNode(nodeName string, namespace string) *corev1.Node {
	node := &corev1.Node{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1alpha4",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      nodeName,
			Namespace: namespace,
		},
		Spec:   corev1.NodeSpec{},
		Status: corev1.NodeStatus{},
	}
	return node
}

func NewCluster(clusterName string, namespace string) *clusterv1.Cluster {
	cluster := &clusterv1.Cluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Cluster",
			APIVersion: "cluster.x-k8s.io/v1alpha4",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterName,
			Namespace: namespace,
		},
		Spec: clusterv1.ClusterSpec{},
	}
	return cluster
}

func NewNamespace(namespace string) *corev1.Namespace {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: namespace},
	}
	return ns
}

func NewSecret(bootstrapSecretName, stringDataValue, namespace string) *corev1.Secret {
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      bootstrapSecretName,
			Namespace: namespace,
		},
		Data: map[string][]byte{
			"value": []byte(stringDataValue),
		},
		Type: "cluster.x-k8s.io/secret",
	}
	return secret
}