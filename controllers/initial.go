package controllers

import (
	"kubeboard/connections"
	"log"
)

func ClusterConnection() {
	_, err := connections.ConnectCluster()
	if err != nil {
		log.Println("Cluster Connection Error:", err.Error())
		return
	}

	log.Println("Cluster Connected")
}
