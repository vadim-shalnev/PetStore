package storeRepository

import (
	"database/sql"
	"github.com/vadim-shalnev/PetStore/models"
)

type Storerepository struct {
	DB *sql.DB
}

type StoreRepository interface {
	NewOrder(order *models.Order) error
	GetOrder(id int) (models.Order, error)
	DeleteOrder(id int) error
	Getinventory(inventory *models.Inventory) error
}
