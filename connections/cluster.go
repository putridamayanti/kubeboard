package connections

import (
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var ActiveConfigFileName = "default"
var KubeClientSet *kubernetes.Clientset
var ApiExtensionClientSet *apiextensionsclientset.Clientset

func ConnectCluster() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", "./configs/"+ActiveConfigFileName)
		if err != nil {
			return nil, err
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	KubeClientSet = clientSet

	apiExtensionClient, err := apiextensionsclientset.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	ApiExtensionClientSet = apiExtensionClient

	return clientSet, nil
}
