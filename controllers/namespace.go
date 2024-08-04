package controllers

import (
	"github.com/gin-gonic/gin"
	v12 "k8s.io/api/core/v1"
	"kubeboard/services"
	"net/http"
)

func GetNamespaces(c *gin.Context) {
	res, err := services.GetNamespaces()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.Items)
}

func GetNamespaceByName(c *gin.Context) {
	namespace := c.Param("namespace")

	res, err := services.GetNamespaceByName(namespace)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateNamespace(c *gin.Context) {
	var request v12.Namespace
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.CreateNamespace(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateNamespace(c *gin.Context) {
	var request v12.Namespace
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.UpdateNamespace(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteNamespace(c *gin.Context) {
	namespace := c.Param("namespace")

	err := services.DeleteNamespace(namespace)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Success")
}
