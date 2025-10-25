package service

import (
	"learn-golang/internal/models"
	"learn-golang/internal/repository"
	"learn-golang/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetAllUser() ([]models.User, error) {
	users, err := us.repo.FindAll()
	if err != nil {
		return nil, utils.WrapError(err, "failed to fetch users", utils.ErrCodeInternal)
	}

	return users, nil
}

func (us *userService) CreateUser(user models.User) (models.User, error) {
	user.Email = utils.NormalizeString(user.Email)

	if _, exist := us.repo.FindByEmail(user.Email); exist {
		return models.User{}, utils.NewError("email already exist", utils.ErrCodeConflict)
	}

	user.UUID = uuid.New().String()

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, utils.WrapError(err, "failed to hash password", utils.ErrCodeInternal)
	}
	user.Password = string(password)

	if err := us.repo.Create(user); err != nil {
		return models.User{}, utils.WrapError(err, "failed to create user", utils.ErrCodeInternal)
	}

	return user, nil
}

func (us *userService) GetUserByUUID(uuid string) (models.User, error) {
	user, found := us.repo.FindByUUID(uuid)
	if !found {
		return models.User{}, utils.NewError("user does not exist", utils.ErrCodeNotFound)
	}

	return user, nil
}

func (us *userService) UpdateUser() {

}

func (us *userService) DeleteUser() {

}
