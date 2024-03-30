package models

type Users []User

type User struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	UserStatus    int    `json:"userStatus"`
	SalesPets     Pets   `json:"salesCount"`
	PurchasesPets Pets   `json:"purchasesCount"`
}

type TokenString struct {
	Token string `json:"auth"`
}
