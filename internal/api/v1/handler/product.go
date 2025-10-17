package handler

import (
	"github.com/gin-gonic/gin"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (u ProductHandler) GetProductsV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "List all products (V1)"})
}

func (u ProductHandler) GetProductsByIdV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Get product by id (V1)"})
}

func (u ProductHandler) GetProductsBySlugV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Get product by slug (V1)"})
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
