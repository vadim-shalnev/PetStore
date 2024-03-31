package controllers

import (
	"github.com/vadim-shalnev/PetStore/internal/Pet/petController"
	"github.com/vadim-shalnev/PetStore/internal/Store/storeController"
	"github.com/vadim-shalnev/PetStore/internal/User/userController"
)

type Controllers struct {
	User  *userController.Usercontroller
	Store *storeController.Storecontroller
	Pet   *petController.Petcontroller
}
