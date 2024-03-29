package storeController

import "net/http"

type Storecontroller struct {
	service StoreService
}

type StoreController interface {
	NewOrder(w http.ResponseWriter, r *http.Request)
	GetOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
	Getinventory(w http.ResponseWriter, r *http.Request)
}
