package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// Status is default method of HealthController
func (h HealthController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
