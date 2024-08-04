package controllers

import (
	"github.com/gin-gonic/gin"
	v12 "k8s.io/api/core/v1"
	"kubeboard/services"
	"net/http"
)

func GetServices(c *gin.Context) {
	namespace := c.Param("namespace")

	res, err := services.GetServices(namespace)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.Items)
}

func GetServiceByName(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	res, err := services.GetServiceByName(namespace, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateService(c *gin.Context) {
	namespace := c.Param("namespace")

	var request v12.Service
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.CreateService(namespace, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateService(c *gin.Context) {
	namespace := c.Param("namespace")

	var request v12.Service
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.UpdateService(namespace, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteService(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	err := services.DeleteService(namespace, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Success")
}
