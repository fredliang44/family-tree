package graphql

import (
	"context"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

// Handler is a func to handler graphql requests
// @Summary Graphql
// @Description Graphql Handler
// @Tags graphql
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Router /graphql [post]
func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Load user
		username := jwt.ExtractClaims(c)["id"]
		ctx := context.WithValue(c.Request.Context(), "User", username)

		// Creates a GraphQL-go HTTP handler with the defined schema
		h := handler.New(&handler.Config{
			Schema: &Schema,
			Pretty: true,
		})

		h.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}
