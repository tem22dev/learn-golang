package repository

import "learn-golang/internal/models"

type InMemoryUserRepository struct {
	users []models.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make([]models.User, 0),
	}
}
