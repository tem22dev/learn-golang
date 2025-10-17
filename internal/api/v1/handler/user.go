package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "List all users (V1)"})
}

func (u UserHandler) GetUsersByIdV1(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"message": "Get user by id (V1)"})
}

func (u UserHandler) GetUsersByUuidV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get user by uuid (V1)",
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
