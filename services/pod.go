package services

import (
	"context"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeboard/connections"
)

func GetPods(namespace string) (*v12.PodList, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetPodByName(namespace string, name string) (*v12.Pod, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreatePod(namespace string, Pod v12.Pod) (*v12.Pod, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Pods(namespace).Create(context.TODO(), &Pod, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdatePod(namespace string, Pod v12.Pod) (*v12.Pod, error) {
	client := connections.KubeClientSet
	res, err := client.CoreV1().Pods(namespace).Update(context.TODO(), &Pod, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeletePod(namespace string, name string) error {
	client := connections.KubeClientSet
	err := client.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
