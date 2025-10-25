package repository

import "learn-golang/internal/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	Create(user models.User) error
	FindByUUID(uuid string) (models.User, bool)
	Update(uuid string, currentUser models.User) error
	Delete()
	FindByEmail(email string) (models.User, bool)
}
