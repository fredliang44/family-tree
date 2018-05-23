package graphql

import (
	"context"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

type contextKey string

var contextKeyUser = contextKey("User")

// Handler is a func to handler graphql requests
func Handler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	return func(c *gin.Context) {
		// Load user
		userID := jwt.ExtractClaims(c)["id"]
		ctx := c.Request.Context()

		h.ServeHTTP(c.Writer, c.Request.WithContext(context.WithValue(ctx, contextKeyUser, userID)))
	}
}
