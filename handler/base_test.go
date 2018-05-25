package handler

import (
	"family-tree/graphql"
	"family-tree/middleware"

	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	r := gofight.New()

	r.GET("/ping").
		SetDebug(true).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, "{\"code\":200,\"message\":\"pong\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

// GinEngine is gin router.
func GinEngine() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	// AUTH & Login
	r.POST("/login", middleware.AuthMiddleware.LoginHandler)

	r.POST("/code", GenCode)
	r.POST("/register_code", GenCode)
	r.POST("/reset_password_code", GenResetCode)

	r.POST("/register", RegisterHandler)
	r.POST("/reset", ResetPassword)

	// HealthCheck
	r.GET("/ping", HealthCheck)

	r.POST("/graphql", graphql.Handler())
	r.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
	r.GET("/init_db", InitDB)

	return r
}
