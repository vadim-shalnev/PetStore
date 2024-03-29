package petController

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vadim-shalnev/PetStore/models"
	"io/ioutil"
	"net/http"
)

func NewPetController(service *PetService) *Petcontroller {
	return &Petcontroller{service: service}
}

func (c *Petcontroller) AddPet(w http.ResponseWriter, r *http.Request) {
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
	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	var status []string
	err = json.Unmarshal(byteBody, &status)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	pets, err := c.service.FindByStatus(status)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, pets)
}

func (c *Petcontroller) GetPet(w http.ResponseWriter, r *http.Request) {
	petId := chi.URLParam(r, "petId")

	pet, err := c.service.GetPet(petId)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, pet)
}

func (c *Petcontroller) ChangePet(w http.ResponseWriter, r *http.Request) {
	petId := chi.URLParam(r, "petId")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	var pet []string
	err = json.Unmarshal(bodyBytes, &pet)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = c.service.ChangePet(petId, pet)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, nil)
}

func (c *Petcontroller) DeletePet(w http.ResponseWriter, r *http.Request) {
	petId := chi.URLParam(r, "petId")
	err := c.service.DeletePet(petId)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, nil)
}
