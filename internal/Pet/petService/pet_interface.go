package petService

import "github.com/vadim-shalnev/PetStore/models"

type Petservice struct {
	repo PetRepository
}

type PetService interface {
	AddPet(pet models.Pet) error
	UpdatePet(pet models.Pet) error
	FindByStatus(status []string) ([]models.Pet, error)
	GetPet(id int) (models.Pet, error)
	ChangePet(id int, pet models.Pet) error
}
