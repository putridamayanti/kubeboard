package controllers

import (
	"github.com/gin-gonic/gin"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"kubeboard/services"
	"net/http"
)

func GetCustomResourceDefinitions(c *gin.Context) {
	res, err := services.GetCustomResourceDefinitions()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.Items)
}

func GetCustomResourceDefinitionByName(c *gin.Context) {
	name := c.Param("name")

	res, err := services.GetCustomResourceDefinitionByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateCustomResourceDefinition(c *gin.Context) {
	var request v1.CustomResourceDefinition
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.CreateCustomResourceDefinition(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateCustomResourceDefinition(c *gin.Context) {
	var request v1.CustomResourceDefinition
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.UpdateCustomResourceDefinition(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteCustomResourceDefinition(c *gin.Context) {
	name := c.Param("name")

	err := services.DeleteCustomResourceDefinition(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Success")
}
