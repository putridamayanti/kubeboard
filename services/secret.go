package services

import (
	"context"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeboard/connections"
)

func GetSecrets(namespace string) (*v12.SecretList, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetSecretByName(namespace string, name string) (*v12.Secret, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateSecret(namespace string, Secret v12.Secret) (*v12.Secret, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Secrets(namespace).Create(context.TODO(), &Secret, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateSecret(namespace string, Secret v12.Secret) (*v12.Secret, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Secrets(namespace).Update(context.TODO(), &Secret, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteSecret(namespace string, name string) error {
	client := connections.KubeClientSet
	err := client.CoreV1().Secrets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
