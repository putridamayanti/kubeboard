package services

import (
	"bytes"
	"context"
	"io"
	v12 "k8s.io/api/core/v1"
	"kubeboard/connections"
)

func GetLogs(namespace string, podName string) (*string, error) {
	client := connections.KubeClientSet
	req := client.CoreV1().Pods(namespace).GetLogs(podName, &v12.PodLogOptions{})
	podLogs, err := req.Stream(context.Background())
	if err != nil {
		return nil, err
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return nil, err
	}
	str := buf.String()

	return &str, nil
}
