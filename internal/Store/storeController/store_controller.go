package storeController

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vadim-shalnev/PetStore/models"
	"io/ioutil"
	"net/http"
)

func NewStoreController(service *StoreService) *Storecontroller {
	return &Storecontroller{
		service: service,
	}
}

func (c *Storecontroller) NewOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	jsonBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	err := json.Unmarshal(jsonBody, &order)
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

	resp, err := c.service.GetOrder(orderID)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}

func (c *Storecontroller) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "orderID")
	err := c.service.DeleteOrder(orderID)
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
