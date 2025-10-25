package handler

import (
	"learn-golang/internal/dto"
	"learn-golang/internal/models"
	"learn-golang/internal/service"
	"learn-golang/internal/utils"
	"learn-golang/internal/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

type GetUserByUuidParam struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

func (uh *UserHandler) GetAllUser(ctx *gin.Context) {
	users, err := uh.service.GetAllUser()
	if err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	usersDTO := dto.MapUsersToDTO(users)

	utils.ResponseSuccess(ctx, http.StatusCreated, &usersDTO)
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	createUser, err := uh.service.CreateUser(user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := dto.MapUserToDTO(createUser)

	utils.ResponseSuccess(ctx, http.StatusCreated, &userDTO)
}

func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {
	var param GetUserByUuidParam
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	user, err := uh.service.GetUserByUUID(param.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := dto.MapUserToDTO(user)
	utils.ResponseSuccess(ctx, http.StatusOK, &userDTO)
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {

}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {

}
