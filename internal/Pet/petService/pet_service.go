package petService

import "github.com/vadim-shalnev/PetStore/models"

func NewPetService(repo petRepository.PetRepository) *Petservice {
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

func (p *Petservice) FindByStatus(status ...string) ([]models.Pet, error) {
	var pets []models.Pet
	for _, s := range status {
		pet := p.repo.FindByStatus(s)
		pets = append(pets, pet)
	}
	return pets, nil
}

func (p *Petservice) GetPet(id int) (models.Pet, error) {
	return p.repo.GetPet(id)
}

func (p *Petservice) ChangePet(id int, name, status string) error {
	var pet models.Pet
	pet.ID = id
	if name != "" {
		pet.Name = name
	}
	if status != "" {
		pet.Status = status
	}
	return p.repo.ChangePet(pet)
}

func (p *Petservice) DeletePet(id int) error {
	return p.repo.DeletePet(id)
}
