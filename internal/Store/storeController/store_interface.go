package storeController

import (
	"github.com/vadim-shalnev/PetStore/internal/Store/storeService"
	"net/http"
)

type Storecontroller struct {
	service storeService.StoreService
}

type StoreController interface {
	NewOrder(w http.ResponseWriter, r *http.Request)
	GetOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
	Getinventory(w http.ResponseWriter, r *http.Request)
}
