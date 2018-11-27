package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Salvatore-Giordano/gin-redis-ip-limiter"
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/fredliang44/family-tree/db"
	"github.com/fredliang44/family-tree/graphql"
	"github.com/fredliang44/family-tree/handler"
	"github.com/fredliang44/family-tree/middleware"
	"github.com/fredliang44/family-tree/utils"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Family Tree API
// @version 0.1
// @description This is a Family Tree API server.

// @contact.name Fred Liang
// @contact.url https://blog.fredliang.cn
// @contact.email info@fredliang.cn

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @license.name MPL-2.0
// @license.url https://github.com/fredliang44/family-tree/blob/master/LICENSE

// @host fmt.fredliang.cn
// @schemes https
func init() {
	if sdn := utils.AppConfig.Sentry.SDN; sdn != "" {
		raven.SetDSN(sdn)
	}
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	// CORS support
	r.Use(middleware.CORSMiddleware())

	// recovery from internal server error
	r.Use(nice.Recovery(utils.RecoveryHandler))

	if gin.Mode() == "release" {
		// Logging to a file.
		f, _ := os.Create("server.log")
		gin.DefaultWriter = io.MultiWriter(f)

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

	files := r.Group("/files")
	{
		files.GET("/token", handler.GetPolicyTokenHandler)
		files.POST("/callback", handler.FilesCallBackHandler)
	}

	// Use ginSwagger gen api doc
	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) { c.Redirect(http.StatusPermanentRedirect, "/doc/index.html") })

	// Show Status
	showStatus()

	// Run Server
	r.Run(utils.AppConfig.Server.Host + ":" + utils.AppConfig.Server.Port)
}

func showStatus() {
	log.Println(
		"\n===================================" +
			"\nAPP         : " + utils.AppConfig.APPName +
			"\nRunning On  : HTTP      " + utils.AppConfig.Server.Host + ":" + utils.AppConfig.Server.Port +
			"\n===================================")
}
