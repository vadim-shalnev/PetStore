package models

import "time"

type Users []User

type User struct {
	ID         int        `json:"id"`
	Username   string     `json:"username"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Phone      string     `json:"phone"`
	UserStatus int        `json:"userStatus"`
	CreatedAt  time.Time  `json:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type TokenString struct {
	Token string `json:"auth"`
}
