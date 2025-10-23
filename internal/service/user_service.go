package service

import (
	"learn-golang/internal/repository"
	"log"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetAllUser() {
	log.Println("Get All User in service")

	us.repo.FindAll()
}

func (us *userService) CreateUser() {

}

func (us *userService) GetUserByUUID() {

}

func (us *userService) UpdateUser() {

}

func (us *userService) DeleteUser() {

}
