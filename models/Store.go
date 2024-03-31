package models

import "time"

type Order struct {
	ID        int        `json:"id"`
	PetID     int        `json:"petId"`
	SellerID  int        `json:"sellerId"`
	BuyerID   int        `json:"buyerId"`
	Quantity  int        `json:"quantity"`
	Status    string     `json:"status"`
	Complete  bool       `json:"complete"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Inventory struct {
	Sold      int `json:"sold"`
	InOrder   int `json:"inStock"`
	Available int `json:"onOrder"`
}
