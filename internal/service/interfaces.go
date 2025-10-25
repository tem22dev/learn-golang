package service

import "learn-golang/internal/models"

type UserService interface {
	GetAllUser(search string, page, limit int) ([]models.User, error)
	GetUserByUUID(uuid string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser()
	DeleteUser()
}
