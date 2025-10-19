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

type PostCategoryV1Param struct {
	Name   string `form:"name" binding:"required"`
	Status string `form:"status" binding:"required,oneof=1 2"`
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetProductsByCategoryV1(ctx *gin.Context) {
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

func (c *CategoryHandler) PostCategoryV1(ctx *gin.Context) {
	var param PostCategoryV1Param
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Post category (V1)",
		"name":    param.Name,
		"status":  param.Status,
	})
}
