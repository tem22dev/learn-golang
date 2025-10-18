package handler

import (
	"learn-golang/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct{}

type GetProductsByCategoryV1Param struct {
	Category string `uri:"category" binding:"oneof=php python golang"`
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c CategoryHandler) GetProductsByCategoryV1(ctx *gin.Context) {
	var param GetProductsByCategoryV1Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Get product by category (V1)",
		"data":    param.Category,
	})
}
