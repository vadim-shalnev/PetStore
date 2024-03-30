package petService

import "github.com/vadim-shalnev/PetStore/models"

func NewPetservice(repo *PetRepository) *Petservice {
	return &Petservice{
		repo: repo,
	}
}

func (p *Petservice) AddPet(pet models.Pet) error {
	return p.repo.AddPet(&pet)
}

func (p *Petservice) UpdatePet(pet models.Pet) error {
	return p.repo.UpdatePet(&pet)
}

func (p *Petservice) FindByStatus(status []string) ([]models.Pet, error) {
	return p.repo.FindByStatus(status)
}

func (p *Petservice) GetPet(id int) (models.Pet, error) {
	return p.repo.GetPet(id)
}
