package storeController

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vadim-shalnev/PetStore/internal/Store/storeService"
	"github.com/vadim-shalnev/PetStore/internal/responder"
	"github.com/vadim-shalnev/PetStore/models"
	"net/http"
	"strconv"
)

func NewStoreController(service storeService.StoreService) *Storecontroller {
	return &Storecontroller{
		service: service,
	}
}

func (c *Storecontroller) NewOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		responder.HandleError(w, err)
		return
	}

	resp, err := c.service.NewOrder(order)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

func (c *Storecontroller) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "orderID")
	ID, err := strconv.Atoi(orderID)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	resp, err := c.service.GetOrder(ID)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

func (c *Storecontroller) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "orderID")
	ID, err := strconv.Atoi(orderID)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err = c.service.DeleteOrder(ID)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, nil)
}

func (c *Storecontroller) Getinventory(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	resp, err := c.service.Getinventory(inventory)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}
