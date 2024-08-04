package controllers

import (
	"github.com/gin-gonic/gin"
	v12 "k8s.io/api/core/v1"
	"kubeboard/services"
	"net/http"
)

func GetPods(c *gin.Context) {
	namespace := c.Param("namespace")

	res, err := services.GetPods(namespace)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.Items)
}

func GetPodByName(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	res, err := services.GetPodByName(namespace, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreatePod(c *gin.Context) {
	namespace := c.Param("namespace")

	var request v12.Pod
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.CreatePod(namespace, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdatePod(c *gin.Context) {
	namespace := c.Param("namespace")

	var request v12.Pod
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.UpdatePod(namespace, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeletePod(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	err := services.DeletePod(namespace, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Success")
}
