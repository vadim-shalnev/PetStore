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
	Sold      int64 `json:"sold"`
	InOrder   int64 `json:"inStock"`
	Available int64 `json:"onOrder"`
}
