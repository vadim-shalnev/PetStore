package petRepository

import (
	"database/sql"
	"github.com/vadim-shalnev/PetStore/models"
	"log"
)

func NewPetRepository(db *sql.DB) *Petrepository {
	return &Petrepository{db}
}

func (r *Petrepository) AddPet(pet *models.Pet) error {
	query := `
	INSERT INTO pets (category_id, owner_id, name, status, created_at, deleted_at)
	VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, NULL)
	`

	_, err := r.DB.Exec(query, pet.Category.ID, pet.OwnerID, pet.Name, pet.Status)
	if err != nil {
		log.Printf("Failed to add pet: %v", err)
		return err
	}

	return nil
}

func (r *Petrepository) UpdatePet(pet *models.Pet) error {
	query := `
	UPDATE pets
	SET category_id = $1, owner_id = $2, name = $3, status = $4, deleted_at = NULL
	WHERE id = $5
	`

	_, err := r.DB.Exec(query, pet.Category.ID, pet.OwnerID, pet.Name, pet.Status, pet.ID)
	if err != nil {
		log.Printf("Failed to update pet: %v", err)
		return err
	}

	return nil
}

func (r *Petrepository) FindByStatus(pets []models.Pet, status string) error {
	query := `
	SELECT id, category_id, owner_id, name, status, created_at, deleted_at
	FROM pets
	WHERE status = $1
	`

	rows, err := r.DB.Query(query, status)
	if err != nil {
		log.Printf("Failed to query pets by status: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var pet models.Pet
		err = rows.Scan(&pet.ID, &pet.Category.ID, &pet.OwnerID, &pet.Name, &pet.Status, &pet.CreatedAt, &pet.DeletedAt)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		pets = append(pets, pet)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return err
	}

	return nil
}

func (r *Petrepository) GetPet(id int) (models.Pet, error) {
	var pet models.Pet
	query := `
	SELECT id, category_id, owner_id, name, status, created_at, deleted_at
	FROM pets
	WHERE id = $1
	`

	err := r.DB.QueryRow(query, id).Scan(&pet.ID, &pet.Category.ID, &pet.OwnerID, &pet.Name, &pet.Status, &pet.CreatedAt, &pet.DeletedAt)
	if err != nil {
		log.Printf("Failed to get pet: %v", err)
		return models.Pet{}, err
	}

	return pet, nil
}

func (r *Petrepository) ChangePet(pet *models.Pet) error {
	query := `
	UPDATE pets
	SET category_id = $1, owner_id = $2, name = $3, status = $4, deleted_at = NULL
	WHERE id = $5
	`

	_, err := r.DB.Exec(query, pet.Category.ID, pet.OwnerID, pet.Name, pet.Status, pet.ID)
	if err != nil {
		log.Printf("Failed to update pet: %v", err)
		return err
	}

	return nil
}

func (r *Petrepository) DeletePet(id int) error {
	query := `
	UPDATE pets
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE id = $1
	`

	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete pet: %v", err)
		return err
	}

	return nil
}
