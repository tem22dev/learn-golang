package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeyMiddleware() gin.HandlerFunc {
	expectedKey := os.Getenv("API_KEY")
	if expectedKey == "" {
		expectedKey = "secret-key"
	}

	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-Api-Key")
		if apiKey == "" {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "Missing X-Api-Key"})
			return
		}

		if apiKey != expectedKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid X-Api-Key"})
			return
		}

		ctx.Next()
	}
}
