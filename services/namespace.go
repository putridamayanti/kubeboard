package services

import (
	"context"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeboard/connections"
)

func GetNamespaces() (*v12.NamespaceList, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetNamespaceByName(namespace string) (*v12.Namespace, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateNamespace(namespace v12.Namespace) (*v12.Namespace, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateNamespace(namespace v12.Namespace) (*v12.Namespace, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Namespaces().Update(context.TODO(), &namespace, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteNamespace(namespace string) error {
	client := connections.KubeClientSet
	err := client.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
