package petService

import (
	"context"
	"github.com/vadim-shalnev/PetStore/internal/Pet/petRepository"
	"github.com/vadim-shalnev/PetStore/internal/middleware"
	"github.com/vadim-shalnev/PetStore/models"
	"strconv"
)

func NewPetService(repo petRepository.PetRepository) *Petservice {
	return &Petservice{
		repo: repo,
	}
}

func (p *Petservice) AddPet(ctx context.Context, pet models.Pet) error {
	token := ctx.Value("jwt_token").(string)
	id, _ := middleware.GetUserinfo(token)
	intID, _ := strconv.Atoi(id)
	pet.OwnerID = intID
	return p.repo.AddPet(&pet)
}

func (p *Petservice) UpdatePet(pet models.Pet) error {
	return p.repo.UpdatePet(&pet)
}

func (p *Petservice) FindByStatus(status string) ([]models.Pet, error) {
	var pets []models.Pet
	err := p.repo.FindByStatus(pets, status)
	if err != nil {
		return nil, err
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
	return p.repo.ChangePet(&pet)
}

func (p *Petservice) DeletePet(id int) error {
	return p.repo.DeletePet(id)
}
