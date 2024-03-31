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

// AddPet @Summary Add a new pet to the store
// @Description get string by ID
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param pet body models.Pet true "Pet object"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /pet/pet [post]
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

// UpdatePet @Summary Update an existing pet
// @Description get string by ID
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param pet body models.Pet true "Pet object"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /pet/pet [put]
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

// FindByStatus @Summary Finds Pets by status
// @Description Only one status for request
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param status query string true "Status values that need to be considered for filter"
// @Success 200 {object} []models.Pet "ok"
// @Failure 400 {string} string "bad request"
// @Router /pet/pet/findByStatus [get]
func (c *Petcontroller) FindByStatus(w http.ResponseWriter, r *http.Request) {
	status := chi.URLParam(r, "status")

	pets, err := c.service.FindByStatus(status)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, pets)
}

// GetPet @Summary Find pet by ID
// @Description Returns a single pet
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param petId path int true "ID of pet"
// @Success 200 {object} models.Pet "ok"
// @Failure 400 {string} string "bad request"
// @Router /pet/pet/{petId} [get]
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

// ChangePet @Summary Change pet
// @Description Change pet
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param petId path int true "ID of pet"
// @Param name query string false "Name of pet"
// @Param status query string false "Status of pet"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /pet/pet/{petId} [post]
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

// DeletePet @Summary Delete pet
// @Description Delete pet
// @Tags Pet
// @Accept  json
// @Produce  json
// @Param petId path int true "ID of pet"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /pet/pet/{petId} [delete]
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
