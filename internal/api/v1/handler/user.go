package handler

import (
	"learn-golang/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

type GetUsersByIdV1Param struct {
	ID int `uri:"id" binding:"gt=0"`
}

type GetUsersByUuidV1Param struct {
	UUID string `uri:"uuid" binding:"uuid"`
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "List all users (V1)"})
}

func (u UserHandler) GetUsersByIdV1(ctx *gin.Context) {
	var param GetUsersByIdV1Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Get user by id (V1)",
		"id":      param.ID,
	})
}

func (u UserHandler) GetUsersByUuidV1(ctx *gin.Context) {
	var param GetUsersByUuidV1Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Get user by uuid (V1)",
		"uuid":    param.UUID,
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
