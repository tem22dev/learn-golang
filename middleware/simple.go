package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// before request
		log.Println("Start func middleware")

		ctx.Next()

		// after request
		log.Println("End func middleware")
	}
}
