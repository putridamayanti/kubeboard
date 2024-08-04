package controllers

import (
	"github.com/gin-gonic/gin"
	"kubeboard/connections"
	"net/http"
)

func SwitchConnection(c *gin.Context) {
	fileName := c.Param("filename")
	connections.ActiveConfigFileName = fileName
	cluster, err := connections.ConnectCluster()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, cluster)
}
