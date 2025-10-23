package service

type UserService interface {
	GetAllUser()
	GetUserByUUID()
	CreateUser()
	UpdateUser()
	DeleteUser()
}
