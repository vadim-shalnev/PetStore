package userService

import (
	"context"
	"github.com/vadim-shalnev/PetStore/internal/User/userRepository"
	"github.com/vadim-shalnev/PetStore/models"
)

type Userservice struct {
	Repository userRepository.UserRepository
}

type UserService interface {
	CreateUser(user models.User) (string, error)
	CreateUsers(users []models.User) ([]string, error)
	Login(user models.User) (string, error)
	Logout(ctx context.Context) (string, error)
	GetUser(username string) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(username string) error
}
