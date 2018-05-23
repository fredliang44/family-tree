package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck is a func
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong", "code": 200})
}
