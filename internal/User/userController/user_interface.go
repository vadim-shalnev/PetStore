package userController

import (
	"github.com/vadim-shalnev/PetStore/internal/User/userService"
	"net/http"
)

type Usercontroller struct {
	service userService.UserService
}

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	CreateUsers(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}
