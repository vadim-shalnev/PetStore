package petService

import (
	"context"
	"github.com/vadim-shalnev/PetStore/internal/Pet/petRepository"
	"github.com/vadim-shalnev/PetStore/models"
)

type Petservice struct {
	repo petRepository.PetRepository
}

type PetService interface {
	AddPet(ctx context.Context, pet models.Pet) error
	UpdatePet(pet models.Pet) error
	FindByStatus(status string) ([]models.Pet, error)
	GetPet(id int) (models.Pet, error)
	ChangePet(id int, name, status string) error
	DeletePet(id int) error
}
