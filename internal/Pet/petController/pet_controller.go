package petController

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vadim-shalnev/PetStore/internal/Pet/petService"
	"github.com/vadim-shalnev/PetStore/internal/responder"
	"github.com/vadim-shalnev/PetStore/models"
	"net/http"
	"strconv"
	"strings"
)

func NewPetController(service petService.PetService) *Petcontroller {
	return &Petcontroller{service: service}
}

func (c *Petcontroller) AddPet(w http.ResponseWriter, r *http.Request) {
	// извлекаем токен из заголовка Authorization
	Usertoken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	ctx := context.WithValue(r.Context(), "jwt_token", Usertoken)
	var pet models.Pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = c.service.AddPet(ctx, pet)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, nil)
}

func (c *Petcontroller) UpdatePet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = c.service.UpdatePet(pet)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, nil)
}

func (c *Petcontroller) FindByStatus(w http.ResponseWriter, r *http.Request) {
	status := chi.URLParam(r, "status")

	pets, err := c.service.FindByStatus(status)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, pets)
}

func (c *Petcontroller) GetPet(w http.ResponseWriter, r *http.Request) {
	petId := chi.URLParam(r, "petId")
	ID, err := strconv.Atoi(petId)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	pet, err := c.service.GetPet(ID)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, pet)
}

func (c *Petcontroller) ChangePet(w http.ResponseWriter, r *http.Request) {
	petId := chi.URLParam(r, "petId")
	ID, err := strconv.Atoi(petId)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	name := r.FormValue("name")
	status := r.FormValue("status")
	err = c.service.ChangePet(ID, name, status)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, nil)
}

func (c *Petcontroller) DeletePet(w http.ResponseWriter, r *http.Request) {
	petId := chi.URLParam(r, "petId")
	ID, err := strconv.Atoi(petId)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = c.service.DeletePet(ID)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, nil)
}
