package user

import "todolist-app/entities"

type User interface {
	GetUserByID(id int) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
	DeleteUser(user entities.User) (entities.User, error)
}
