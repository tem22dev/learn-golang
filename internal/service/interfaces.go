package service

import "learn-golang/internal/models"

type UserService interface {
	GetAllUser()
	GetUserByUUID()
	CreateUser(user models.User) (models.User, error)
	UpdateUser()
	DeleteUser()
}
