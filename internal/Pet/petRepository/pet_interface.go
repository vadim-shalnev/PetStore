package petRepository

import (
	"database/sql"
	"github.com/vadim-shalnev/PetStore/models"
)

type Petrepository struct {
	DB *sql.DB
}

type PetRepository interface {
	AddPet(pet *models.Pet) error
	UpdatePet(pet *models.Pet) error
	FindByStatus(pets []models.Pet, status string) error
	GetPet(id int) (models.Pet, error)
	ChangePet(pet *models.Pet) error
	DeletePet(id int) error
}
