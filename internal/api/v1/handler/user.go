package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "List all users (V1)"})
}

func (u UserHandler) GetUsersByIdV1(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "id must be integer"})
		return
	}

	if id < 0 {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "id must be >= 0"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get user by id (V1)", "id": id})
}

func (u UserHandler) GetUsersByUuidV1(ctx *gin.Context) {
	idStr := ctx.Param("uuid")
	_, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "id must be a valid uuid"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Get user by uuid (V1)",
		"uuid":    idStr,
	})
}

func (u UserHandler) PostUsersV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Create user (V1)"})
}

func (u UserHandler) PutUsersV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Update user (V1)"})
}

func (u UserHandler) DeleteUsersV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Delete user (V1)"})
}
