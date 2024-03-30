package petService

import "github.com/vadim-shalnev/PetStore/models"

type Petservice struct {
	repo petRepository.PetRepository
}

type PetService interface {
	AddPet(pet models.Pet) error
	UpdatePet(pet models.Pet) error
	FindByStatus(status ...string) ([]models.Pet, error)
	GetPet(id int) (models.Pet, error)
	ChangePet(id int, name, status string) error
	DeletePet(id int) error
}
