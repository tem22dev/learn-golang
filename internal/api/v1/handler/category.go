package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct{}

var validCategory = map[string]bool{
	"php":    true,
	"python": true,
	"golang": true,
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c CategoryHandler) GetProductsByCategoryV1(ctx *gin.Context) {
	category := ctx.Param("category")

	if !validCategory[category] {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "invalid category"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Get product by category (V1)", "data": category})
}
