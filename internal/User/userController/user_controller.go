package userController

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
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
	resp, err := c.service.CreateUser(User)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
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
	resp, err := c.service.CreateUser(Users)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

func (c *Usercontroller) Login(w http.ResponseWriter, r *http.Request) {
	Usertoken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	ctx := context.WithValue(r.Context(), "jwt_token", Usertoken)

	resp, err := c.service.Login(ctx)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
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
	userID := chi.URLParam(r, "username")

	resp, err := c.service.GetUser(ctx, userID)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

func (c *Usercontroller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	Usertoken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	ctx := context.WithValue(r.Context(), "jwt_token", Usertoken)
	userID := chi.URLParam(r, "username")

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
	resp, err := c.service.UpdateUser(ctx, userID, User)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

func (c *Usercontroller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "username")
	resp, err := c.service.DeleteUser(userID)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}
