package models

type Order struct {
	ID       int64  `json:"id"`
	PetID    int64  `json:"petId"`
	Quantity int64  `json:"quantity"`
	ShipDate string `json:"shipDate"`
	Status   string `json:"status"`
	Complete bool   `json:"complete"`
}

type Inventory struct {
	Sold      int64 `json:"sold"`
	InOrder   int64 `json:"inStock"`
	Available int64 `json:"onOrder"`
}
