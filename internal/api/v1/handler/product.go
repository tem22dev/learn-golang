package handler

import (
	"learn-golang/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{}

type GetProductsBySlugV1Param struct {
	Slug string `uri:"slug" binding:"slug,min=5,max=50"`
}

type GetProductsV1Param struct {
	Search string `form:"search" binding:"required,min=3,max=50,search"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (u ProductHandler) GetProductsV1(ctx *gin.Context) {
	var param GetProductsV1Param
	if err := ctx.ShouldBindQuery(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	if param.Limit == 0 {
		param.Limit = 10
	}

	ctx.JSON(200, gin.H{
		"message": "List all products (V1)",
		"search":  param.Search,
		"limit":   param.Limit,
	})
}

func (u ProductHandler) GetProductsByIdV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Get product by id (V1)"})
}

func (u ProductHandler) GetProductsBySlugV1(ctx *gin.Context) {
	var param GetProductsBySlugV1Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Get product by slug (V1)",
		"slug":    param.Slug,
	})
}

func (u ProductHandler) PostProductsV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Create product (V1)"})
}

func (u ProductHandler) PutProductsV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Update product (V1)"})
}

func (u ProductHandler) DeleteProductsV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Delete product (V1)"})
}
