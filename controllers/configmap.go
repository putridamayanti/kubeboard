package controllers

import (
	"github.com/gin-gonic/gin"
	v12 "k8s.io/api/core/v1"
	"kubeboard/services"
	"net/http"
)

func GetConfigMaps(c *gin.Context) {
	namespace := c.Param("namespace")

	res, err := services.GetConfigMaps(namespace)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.Items)
}

func GetConfigMapByName(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	res, err := services.GetConfigMapByName(namespace, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateConfigMap(c *gin.Context) {
	namespace := c.Param("namespace")

	var request v12.ConfigMap
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.CreateConfigMap(namespace, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateConfigMap(c *gin.Context) {
	namespace := c.Param("namespace")

	var request v12.ConfigMap
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.UpdateConfigMap(namespace, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteConfigMap(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	err := services.DeleteConfigMap(namespace, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Success")
}
