package userRepository

import (
	"database/sql"
	"github.com/vadim-shalnev/PetStore/models"
)

type Userrepository struct {
	DB *sql.DB
}

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(usernaame string) error
}
