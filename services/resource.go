package services

import (
	"context"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeboard/connections"
)

func GetCustomResourceDefinitions() (*v1.CustomResourceDefinitionList, error) {
	client := connections.ApiExtensionClientSet
	res, err := client.ApiextensionsV1().CustomResourceDefinitions().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetCustomResourceDefinitionByName(name string) (*v1.CustomResourceDefinition, error) {
	client := connections.ApiExtensionClientSet
	res, err := client.ApiextensionsV1().CustomResourceDefinitions().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateCustomResourceDefinition(params v1.CustomResourceDefinition) (*v1.CustomResourceDefinition, error) {
	client := connections.ApiExtensionClientSet
	res, err := client.ApiextensionsV1().CustomResourceDefinitions().Create(context.TODO(), &params, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateCustomResourceDefinition(params v1.CustomResourceDefinition) (*v1.CustomResourceDefinition, error) {
	client := connections.ApiExtensionClientSet
	res, err := client.ApiextensionsV1().CustomResourceDefinitions().Update(context.TODO(), &params, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteCustomResourceDefinition(name string) error {
	client := connections.ApiExtensionClientSet
	err := client.ApiextensionsV1().CustomResourceDefinitions().Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
