package handler

import (
	"family-tree/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary HealthCheck
// @Description Respond pong to ping request
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Ping
// @Router /ping [get]
// HealthCheck is a func
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Ping{Code: http.StatusOK, Message: "pong"})
}
