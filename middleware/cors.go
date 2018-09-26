package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware is a func to handler CORS request
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if gin.Mode() == "release" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
