package controllers

import (
	"github.com/gin-gonic/gin"
	v12 "k8s.io/api/core/v1"
	"kubeboard/services"
	"net/http"
)

func GetSecrets(c *gin.Context) {
	namespace := c.Param("namespace")

	res, err := services.GetSecrets(namespace)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.Items)
}

func GetSecretByName(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	res, err := services.GetSecretByName(namespace, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateSecret(c *gin.Context) {
	namespace := c.Param("namespace")

	var request v12.Secret
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.CreateSecret(namespace, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateSecret(c *gin.Context) {
	namespace := c.Param("namespace")

	var request v12.Secret
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.UpdateSecret(namespace, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteSecret(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	err := services.DeleteSecret(namespace, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Success")
}
