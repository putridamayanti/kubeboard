package controllers

import (
	"github.com/gin-gonic/gin"
	"kubeboard/services"
	"net/http"
)

func GetNodes(c *gin.Context) {
	res, err := services.GetListNodes()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res.Items)
}

func GetNodeByName(c *gin.Context) {
	name := c.Param("name")

	res, err := services.GetNodeByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
