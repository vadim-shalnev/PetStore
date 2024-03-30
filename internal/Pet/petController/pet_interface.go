package petController

import (
	"github.com/vadim-shalnev/PetStore/internal/Pet/petService"
	"net/http"
)

type Petcontroller struct {
	service petService.PetService
}

type PetController interface {
	AddPet(w http.ResponseWriter, r *http.Request)
	UpdatePet(w http.ResponseWriter, r *http.Request)
	FindByStatus(w http.ResponseWriter, r *http.Request)
	GetPet(w http.ResponseWriter, r *http.Request)
	ChangePet(w http.ResponseWriter, r *http.Request)
	DeletePet(w http.ResponseWriter, r *http.Request)
}
