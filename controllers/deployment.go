package controllers

import (
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	"kubeboard/services"
	"net/http"
)

func GetDeployments(c *gin.Context) {
	namespace := c.Param("namespace")

	res, err := services.GetDeployments(namespace)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.Items)
}

func GetDeploymentByName(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	res, err := services.GetDeploymentByName(namespace, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateDeployment(c *gin.Context) {
	namespace := c.Param("namespace")

	var request v1.Deployment
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.CreateDeployment(namespace, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateDeployment(c *gin.Context) {
	namespace := c.Param("namespace")

	var request v1.Deployment
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.UpdateDeployment(namespace, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteDeployment(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	res, err := services.DeleteDeployment(namespace, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
