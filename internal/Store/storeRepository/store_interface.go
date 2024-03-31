package storeRepository

import (
	"database/sql"
	"errors"
	"github.com/vadim-shalnev/PetStore/models"
	"log"
)

func NewStoreRepository(db *sql.DB) *Storerepository {
	return &Storerepository{DB: db}
}

func (r *Storerepository) NewOrder(order *models.Order) error {
	query := `
	INSERT INTO orders (pet_id, seller_id, buyer_id, quantity, status, complete, created_at, deleted_at)
	VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, NULL)
	`

	_, err := r.DB.Exec(query, order.PetID, order.SellerID, order.BuyerID, order.Quantity, order.Status, order.Complete)
	if err != nil {
		log.Printf("Failed to create new order: %v", err)
		return err
	}

	return nil
}

func (r *Storerepository) GetOrder(id int) (models.Order, error) {
	var order models.Order

	query := `
	SELECT id, pet_id, seller_id, buyer_id, quantity, status, complete, created_at, deleted_at
	FROM orders
	WHERE id = $1
	`

	err := r.DB.QueryRow(query, id).Scan(&order.ID, &order.PetID, &order.SellerID, &order.BuyerID, &order.Quantity, &order.Status, &order.Complete, &order.CreatedAt, &order.DeletedAt)
	if err != nil {
		log.Printf("Failed to get order: %v", err)
		return order, err
	}
	if order.DeletedAt != nil {
		return order, errors.New("order not found")
	}

	return order, nil
}

func (r *Storerepository) DeleteOrder(id int) error {
	query := `
	UPDATE orders
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE id = $1
	`

	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete order: %v", err)
		return err
	}

	return nil
}

func (r *Storerepository) Getinventory(inventory *models.Inventory) error {
	query := `
		SELECT 
			(SELECT COUNT(*) FROM orders WHERE status = 'Available') AS available,
			(SELECT COUNT(*) FROM orders WHERE status = 'InOrder') AS in_order,
			(SELECT COUNT(*) FROM orders WHERE status = 'Sold') AS sold
	`

	err := r.DB.QueryRow(query).Scan(&inventory.Available, &inventory.InOrder, &inventory.Sold)
	if err != nil {
		log.Printf("Failed to get inventory status: %v", err)
		return err
	}

	return nil
}
