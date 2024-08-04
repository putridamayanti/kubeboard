package services

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeboard/connections"
)

func GetDeployments(namespace string) (*v1.DeploymentList, error) {
	client := connections.KubeClientSet
	res, err := client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetDeploymentByName(namespace string, name string) (*v1.Deployment, error) {
	client := connections.KubeClientSet
	res, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateDeployment(namespace string, deployment v1.Deployment) (*v1.Deployment, error) {
	client := connections.KubeClientSet
	res, err := client.AppsV1().Deployments(namespace).Create(context.TODO(), &deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateDeployment(namespace string, deployment v1.Deployment) (*v1.Deployment, error) {
	client := connections.KubeClientSet
	res, err := client.AppsV1().Deployments(namespace).Update(context.TODO(), &deployment, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteDeployment(namespace string, name string) error {
	client := connections.KubeClientSet
	err := client.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
