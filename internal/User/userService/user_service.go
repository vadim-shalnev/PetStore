package userService

import (
	"context"
	"fmt"
	"github.com/vadim-shalnev/PetStore/internal/criptografy"
	"github.com/vadim-shalnev/PetStore/internal/middleware"
	"github.com/vadim-shalnev/PetStore/models"
)

func NewUserService(repo userService.UserRepos) *Userservice {
	return &Userservice{repo: repo}
}

func (u *Userservice) CreateUser(user models.User) (string, error) {
	//хэшируем пароль
	err := criptografy.HashPassword(&user)
	if err != nil {
		return "", err
	}
	// добавляем в бд
	err = u.repo.CreateUser(&user)
	if err != nil {
		return "", err
	}
	// генерируем токен
	token, err := middleware.NewTokenMiddleware(&user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *Userservice) CreateUsers(users []models.User) ([]string, error) {
	var resp []string
	for _, user := range users {
		err := criptografy.HashPassword(&user)
		if err != nil {
			resp = append(resp, fmt.Sprintf("internalError with user %s: %s", user.Username, err))
		}
		// добавляем в бд
		err = u.repo.CreateUser(&user)
		if err != nil {
			resp = append(resp, fmt.Sprintf("internalError with user %s: %s", user.Username, err))
		}
		// генерируем токен
		token, err := middleware.NewTokenMiddleware(&user)
		if err != nil {
			resp = append(resp, fmt.Sprintf("internalError with user %s: %s", user.Username, err))
		}
		resp = append(resp, fmt.Sprintf("User username %s create with ID: %v token is %s", user.Username, user.ID, token))
	}
	return resp, nil
}

func (u *Userservice) Login(user models.User) (string, error) {
	// ищем пользователя в бд
	password := user.Password
	err := u.repo.GetUserBy_username(&user)
	if err != nil {
		return "", err
	}
	// проверяем пароль
	err = criptografy.CheckPassword(user.Password, password)
	if err != nil {
		return "", err
	}
	// генерируем токен
	token, err := middleware.NewTokenMiddleware(&user)
	if err != nil {
		return "", err
	}
	return token, nil
}

// удаляем токен клиента
func (u *Userservice) Logout(ctx context.Context) (string, error) {
	return "", nil
}

func (u *Userservice) GetUser(username string) (models.User, error) {
	var user models.User
	user.Username = username
	err := u.repo.GetUserBy_userID(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *Userservice) UpdateUser(user models.User, userName string) (models.User, error) {
	user, err := u.repo.UpdateUser(&user, userName)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (u *Userservice) DeleteUser(username string) error {
	return u.repo.DeleteUser(username)
}
