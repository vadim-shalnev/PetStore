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

// NewOrder @Summary Create new order
// @Description Create new order
// @Tags Store
// @Accept json
// @Produce json
// @Param order body models.Order true "Order"
// @Success 200 {object} models.Order
// @Failure 400 {object} string "bad request"
// @Router /store/order [post]
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

// GetOrder @Summary Get order by ID
// @Description Get order by ID
// @Tags Store
// @Accept json
// @Produce json
// @Param orderID path int true "Order ID"
// @Success 200 {object} models.Order
// @Failure 400 {object} string "bad request"
// @Router /store/order/{orderID} [get]
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

// DeleteOrder @Summary Delete order by ID
// @Description Delete order by ID
// @Tags Store
// @Accept json
// @Produce json
// @Param orderID path int true "Order ID"
// @Success 200 {object} string
// @Failure 400 {object} string "bad request"
// @Router /store/order/{orderID} [delete]
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

// Getinventory @Summary Get inventory
// @Description Get inventory
// @Tags Store
// @Accept json
// @Produce json
// @Success 200 {object} models.Inventory
// @Failure 400 {object} string "bad request"
// @Router /api/store/inventory [get]
func (c *Storecontroller) Getinventory(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	resp, err := c.service.Getinventory(inventory)
	if err != nil {
		responder.HandleError(w, err)
		return
	}
	responder.HandleSuccess(w, resp)
}
