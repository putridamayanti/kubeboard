# Kubeboard - Dashboard API for Kubernetes

This is api for kubernetes dashboard. You can manage deployment, service, pod, etc.

## Install

    go get .

## Run the app

    go run main.go


# REST API

The REST API is described below.

## Check Health
* `GET` `/health` : Get cluster health status


## Cluster
* `GET` `/switch-cluster/:filename` : Switch cluster to another kubeconfig filename in `configs`

## Namespace
* `GET` `/namespace` : Get all namespace
* `GET` `/namespace/:name` : Get a namespace by name
* `POST` `/namespace` : Create namespace
* `PATCH` `/namespace` : Update namespace
* `DELETE` `/namespace/:name` : Delete namespace by name

## Node
* `GET` `/node` : Get all node
* `GET` `/node/:name` : Get a node by name

## Custom Resource Definition
* `GET` `/crd` : Get all crd
* `GET` `/crd/:name` : Get a crd by name
* `POST` `/crd` : Create crd
* `PATCH` `/crd` : Update crd
* `DELETE` `/crd/:name` : Delete crd by name

## Config Map
* `GET` `/config-map` : Get all config-map
* `GET` `/config-map/:name` : Get a config-map by name
* `POST` `/config-map` : Create config-map
* `PATCH` `/config-map` : Update config-map
* `DELETE` `/config-map/:name` : Delete config-map by name

## Deployment
* `GET` `/deployment` : Get all deployment
* `GET` `/deployment/:name` : Get a deployment by name
* `POST` `/deployment` : Create deployment
* `PATCH` `/deployment` : Update deployment
* `DELETE` `/deployment/:name` : Delete deployment by name

## Pod
* `GET` `/pod` : Get all pod
* `GET` `/pod/:name` : Get a pod by name
* `POST` `/pod` : Create pod
* `PATCH` `/pod` : Update pod
* `DELETE` `/pod/:name` : Delete pod by name

## Secret
* `GET` `/secret` : Get all secret
* `GET` `/secret/:name` : Get a secret by name
* `POST` `/secret` : Create secret
* `PATCH` `/secret` : Update secret
* `DELETE` `/secret/:name` : Delete secret by name

## Service
* `GET` `/service` : Get all service
* `GET` `/service/:name` : Get a service by name
* `POST` `/service` : Create service
* `PATCH` `/service` : Update service
* `DELETE` `/service/:name` : Delete service by name