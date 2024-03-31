package storeService

import (
	"github.com/vadim-shalnev/PetStore/internal/Store/storeRepository"
	"github.com/vadim-shalnev/PetStore/models"
)

type Storeservice struct {
	repo storeRepository.StoreRepository
}

type StoreService interface {
	NewOrder(order models.Order) (models.Order, error)
	GetOrder(id int) (models.Order, error)
	DeleteOrder(id int) error
	Getinventory(inventory models.Inventory) (models.Inventory, error)
}
