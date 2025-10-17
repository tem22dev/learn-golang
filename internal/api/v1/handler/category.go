package handler

import (
	"learn-golang/utils"
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
	if err := utils.ValidationIntList("Category", category, validCategory); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get product by category (V1)", "data": category})
}
