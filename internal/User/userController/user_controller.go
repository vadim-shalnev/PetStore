package userController

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vadim-shalnev/PetStore/internal/responder"
	"github.com/vadim-shalnev/PetStore/models"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewUserController(service *userService) *Usercontroller {
	return &Usercontroller{service: service}
}

func (c *Usercontroller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var User models.User
	jsonBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = json.Unmarshal(jsonBody, &User)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = c.service.CreateUser(User)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, nil)
}

func (c *Usercontroller) CreateUsers(w http.ResponseWriter, r *http.Request) {
	var Users []models.User
	jsonBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = json.Unmarshal(jsonBody, &Users)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	resp, err := c.service.CreateUsers(Users)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

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

func (c *Usercontroller) Logout(w http.ResponseWriter, r *http.Request) {
	Usertoken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	ctx := context.WithValue(r.Context(), "jwt_token", Usertoken)

	resp, err := c.service.Logout(ctx)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

func (c *Usercontroller) GetUser(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "username")

	resp, err := c.service.GetUser(userName)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

func (c *Usercontroller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "username")

	var User models.User
	jsonBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = json.Unmarshal(jsonBody, &User)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	resp, err := c.service.UpdateUser(userName, User)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

func (c *Usercontroller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "username")
	resp, err := c.service.DeleteUser(userName)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}
