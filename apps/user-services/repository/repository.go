package repository

type UserRepository interface {
	Open()
	CreateUser()
	GetUserById()
	GetUserByUsername()
	GetUserByEmail()
	GetUsers()
	UpdateUser()
	DeleteUser()
	Close()
}
