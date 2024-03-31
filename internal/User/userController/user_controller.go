package userController

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vadim-shalnev/PetStore/internal/User/userService"
	"github.com/vadim-shalnev/PetStore/internal/responder"
	"github.com/vadim-shalnev/PetStore/models"
	"net/http"
	"strings"
)

func NewUserController(service userService.UserService) *Usercontroller {
	return &Usercontroller{service: service}
}

// CreateUser @Summary CreateUser
// @Description CreateUser
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "user"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /user/user [post]
func (c *Usercontroller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	token, err := c.service.CreateUser(user)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	w.Header().Set("Authorization", "Bearer "+token)
	responder.HandleSuccess(w, token)
}

// CreateUsers @Summary CreateUsers
// @Description CreateUsers
// @Tags User
// @Accept json
// @Produce json
// @Param users body []models.User true "users"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /user/user/createWithList [post]
func (c *Usercontroller) CreateUsers(w http.ResponseWriter, r *http.Request) {
	var Users []models.User
	err := json.NewDecoder(r.Body).Decode(&Users)
	resp, err := c.service.CreateUsers(Users)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

// Login @Summary Login
// @Description Login
// @Tags User
// @Accept json
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /user/user/login [post]
func (c *Usercontroller) Login(w http.ResponseWriter, r *http.Request) {
	// Извлекаем из запроса username и password и контекст для логина.
	login := r.FormValue("Username")
	Password := r.FormValue("Password")
	var user models.User
	user.Username = login
	user.Password = Password
	resp, err := c.service.Login(user)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	w.Header().Set("Authorization", "Bearer "+resp)
	responder.HandleSuccess(w, nil)
}

// Logout @Summary Logout
// @Description Logout
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /user/user/logout [get]
func (c *Usercontroller) Logout(w http.ResponseWriter, r *http.Request) {
	Usertoken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	ctx := context.WithValue(r.Context(), "jwt_token", Usertoken)

	resp, err := c.service.Logout(ctx)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	w.Header().Set("Authorization", "Bearer "+resp)
	responder.HandleSuccess(w, resp)
}

// GetUser @Summary GetUser
// @Description GetUser
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Success 200 {object} models.User "ok"
// @Failure 400 {string} string "bad request"
// @Router /user/user [get]
func (c *Usercontroller) GetUser(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "username")

	resp, err := c.service.GetUser(userName)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

// UpdateUser @Summary UpdateUser
// @Description UpdateUser
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Param user body models.User true "user"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /user/user/{username} [put]
func (c *Usercontroller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = c.service.UpdateUser(user)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, nil)
}

// DeleteUser @Summary DeleteUser
// @Description DeleteUser
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /user/user/{username} [delete]
func (c *Usercontroller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "username")
	err := c.service.DeleteUser(userName)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, nil)
}
