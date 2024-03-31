package models

import "time"

type Pet struct {
	ID        int        `json:"id"`
	Category  Category   `json:"category"`
	OwnerID   int        `json:"owner_id"`
	Name      string     `json:"name"`
	Tags      []Category `json:"tags"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
