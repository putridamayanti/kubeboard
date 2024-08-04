package services

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"kubeboard/connections"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetListNodes() (*v1.NodeList, error) {
	client := connections.KubeClientSet
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func GetNodeByName(name string) (*v1.Node, error) {
	client := connections.KubeClientSet
	node, err := client.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return node, nil
}
