package services

import (
	"context"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeboard/connections"
)

func GetServices(namespace string) (*v12.ServiceList, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetServiceByName(namespace string, name string) (*v12.Service, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateService(namespace string, Service v12.Service) (*v12.Service, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Services(namespace).Create(context.TODO(), &Service, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateService(namespace string, Service v12.Service) (*v12.Service, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Services(namespace).Update(context.TODO(), &Service, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteService(namespace string, name string) error {
	client := connections.KubeClientSet
	err := client.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
