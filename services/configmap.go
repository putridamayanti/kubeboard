package services

import (
	"context"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeboard/connections"
)

func GetConfigMaps(namespace string) (*v12.ConfigMapList, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetConfigMapByName(namespace string, name string) (*v12.ConfigMap, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateConfigMap(namespace string, ConfigMap v12.ConfigMap) (*v12.ConfigMap, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().ConfigMaps(namespace).Create(context.TODO(), &ConfigMap, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateConfigMap(namespace string, ConfigMap v12.ConfigMap) (*v12.ConfigMap, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().ConfigMaps(namespace).Update(context.TODO(), &ConfigMap, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteConfigMap(namespace string, name string) error {
	client := connections.KubeClientSet
	err := client.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
