package repository

import (
	"learn-golang/internal/models"
	"log"
)

type InMemoryUserRepository struct {
	users []models.User
}

func NewInMemoryUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users: make([]models.User, 0),
	}
}

func (ur *InMemoryUserRepository) FindAll() {
	log.Println("Get All User in repository")
}

func (ur *InMemoryUserRepository) Create() {

}

func (ur *InMemoryUserRepository) FindByUUID() {

}

func (ur *InMemoryUserRepository) Update() {

}

func (ur *InMemoryUserRepository) Delete() {

}
