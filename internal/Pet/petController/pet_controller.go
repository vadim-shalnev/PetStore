package petController

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/vadim-shalnev/PetStore/internal/Pet/petService"
	"github.com/vadim-shalnev/PetStore/internal/responder"
	"github.com/vadim-shalnev/PetStore/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func NewPetController(service petService.PetService) *Petcontroller {
	return &Petcontroller{service: service}
}

func (c *Petcontroller) AddPet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = c.service.AddPet(pet)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, nil)
}

func (c *Petcontroller) UpdatePet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	jsonBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = json.Unmarshal(jsonBody, &pet)
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
	statuses := r.URL.Query()["status"]
	if len(statuses) == 0 {
		responder.HandleError(w, errors.New("No status provided"))
		return
	}

	pets, err := c.service.FindByStatus(statuses...)
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
