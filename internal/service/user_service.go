package service

import (
	"learn-golang/internal/models"
	"learn-golang/internal/repository"
	"learn-golang/internal/utils"
	"strings"

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

func (us *userService) GetAllUser(search string, page, limit int) ([]models.User, error) {
	users, err := us.repo.FindAll()
	if err != nil {
		return nil, utils.WrapError(err, "failed to fetch users", utils.ErrCodeInternal)
	}

	var filteredUsers []models.User
	if search != "" {
		search = strings.ToLower(search)
		for _, user := range users {
			name := strings.ToLower(user.Name)
			email := strings.ToLower(user.Email)

			if strings.Contains(name, search) || strings.Contains(email, search) {
				filteredUsers = append(filteredUsers, user)
			}
		}
	} else {
		filteredUsers = users
	}

	start := (page - 1) * limit
	if start >= len(filteredUsers) {
		return []models.User{}, nil
	}

	end := start + limit
	if end >= len(filteredUsers) {
		end = len(filteredUsers)
	}

	return filteredUsers[start:end], nil
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

func (us *userService) UpdateUser(uuid string, updateUser models.User) (models.User, error) {
	updateUser.Email = utils.NormalizeString(updateUser.Email)

	if _, exist := us.repo.FindByEmail(updateUser.Email); exist {
		return models.User{}, utils.NewError("email already exist", utils.ErrCodeConflict)
	}

	currentUser, found := us.repo.FindByUUID(uuid)
	if !found {
		return models.User{}, utils.NewError("user does not exist", utils.ErrCodeNotFound)
	}

	currentUser.Name = updateUser.Name
	currentUser.Email = updateUser.Email
	currentUser.Age = updateUser.Age
	currentUser.Level = updateUser.Level
	currentUser.Status = updateUser.Status

	if updateUser.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, utils.WrapError(err, "failed to hash password", utils.ErrCodeInternal)
		}
		currentUser.Password = string(password)
	}

	if err := us.repo.Update(uuid, currentUser); err != nil {
		return models.User{}, utils.WrapError(err, "failed to update user", utils.ErrCodeInternal)
	}

	return currentUser, nil
}

func (us *userService) DeleteUser() {

}
