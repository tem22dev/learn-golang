package repository

import "learn-golang/internal/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	Create(user models.User) error
	FindByUUID(uuid string) (models.User, bool)
	Update()
	Delete()
	FindByEmail(email string) (models.User, bool)
}
