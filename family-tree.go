package main

import (
	"family-tree/db"
	"family-tree/graphql"
	"family-tree/handler"
	"family-tree/middleware"
	"family-tree/utils"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Salvatore-Giordano/gin-redis-ip-limiter"
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
)

func init() {
	if sdn := utils.AppConfig.Sentry.SDN; sdn != "" {
		raven.SetDSN(sdn)
	}
}

func main() {
	r := gin.Default()

	// CORS support
	r.Use(middleware.CORSMiddleware())

	if gin.Mode() == "release" {
		// Logging to a file.
		f, _ := os.Create("server.log")
		gin.DefaultWriter = io.MultiWriter(f)

		// recovery from internal server error
		r.Use(nice.Recovery(utils.RecoveryHandler))

		// limit request frequency per minute
		r.Use(iplimiter.NewRateLimiterMiddleware(db.RedisClient, "general", 200, time.Minute))
	}

	// AUTH & Login
	r.POST("/login", middleware.AuthMiddleware.LoginHandler)

	r.POST("/code", handler.GenCode)
	r.POST("/register_code", handler.GenCode)
	r.POST("/reset_password_code", handler.GenResetCode)

	r.POST("/register", handler.RegisterHandler)
	r.POST("/reset", handler.ResetPassword)

	// HealthCheck
	r.GET("/ping", handler.HealthCheck)

	// Main Handler
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		auth.POST("/graphql", graphql.Handler())
		auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
		auth.GET("/init_db", handler.InitDB)
	}

	// Show Status
	showStatus()

	// Run Server
	r.Run(utils.AppConfig.Server.Host + ":" + utils.AppConfig.Server.Port)
}

func showStatus() {
	fmt.Println(
		"\n===================================" +
			"\nAPP         : " + utils.AppConfig.APPName +
			"\nRunning On  : HTTP      " + utils.AppConfig.Server.Host + ":" + utils.AppConfig.Server.Port +
			"\n===================================")
}
