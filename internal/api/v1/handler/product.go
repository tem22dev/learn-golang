package handler

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{}

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (u ProductHandler) GetProductsV1(ctx *gin.Context) {
	search := ctx.Query("search")
	if search == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "query parameter search required"})
		return
	}

	if len(search) < 3 || len(search) > 50 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "query parameter search too long"})
		return
	}

	if !searchRegex.MatchString(search) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "query parameter search invalid"})
		return
	}

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "query parameter limit invalid"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "List all products (V1)",
		"search":  search,
		"limit":   limit,
	})
}

func (u ProductHandler) GetProductsByIdV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Get product by id (V1)"})
}

func (u ProductHandler) GetProductsBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if !slugRegex.MatchString(slug) {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "slug is not valid"})
		return
	}
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
