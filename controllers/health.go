package controllers

import (
	"github.com/gin-gonic/gin"
	"kubeboard/services"
	"net/http"
)

func CheckHealth(c *gin.Context) {
	nodes, err := services.GetListNodes()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nodes)
}
