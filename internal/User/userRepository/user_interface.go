package userRepository

import (
	"database/sql"
	"errors"
	"github.com/vadim-shalnev/PetStore/models"
	"log"
)

func NewUserRepository(db *sql.DB) *Userrepository {
	return &Userrepository{
		DB: db,
	}
}

func (r *Userrepository) CreateUser(user *models.User) error {
	query := `
	INSERT INTO users (username, first_name, last_name, email, password, phone, user_status, created_at, deleted_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, NULL)
	`

	_, err := r.DB.Exec(query, user.Username, user.FirstName, user.LastName, user.Email, user.Password, user.Phone, user.UserStatus)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return err
	}
	return nil
}

func (r *Userrepository) GetUserByUsername(user *models.User) error {
	query := `
	SELECT id, first_name, last_name, email, password, phone, user_status, created_at, deleted_at
	FROM users
	WHERE username = $1
	`

	err := r.DB.QueryRow(query, user.Username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Phone, &user.UserStatus, &user.CreatedAt, &user.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("User not found")
		}
		log.Printf("Failed to get user by username: %v", err)
		return err
	}
	// проверяем что пользователь не удален
	if user.DeletedAt != nil {
		return errors.New("User is deleted")
	}
	return nil
}

func (r *Userrepository) UpdateUser(user *models.User) error {
	query := `
	UPDATE users
	SET first_name = $1, last_name = $2, email = $3, password = $4, phone = $5, user_status = $6, created_at = $7, deleted_at = $8
	WHERE username = $9
	`
	_, err := r.DB.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.Phone, user.UserStatus, user.CreatedAt, user.DeletedAt, user.Username)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return err
	}
	// проверяем что пользователь не удален
	if user.DeletedAt != nil {
		return errors.New("User is deleted")
	}
	return nil
}

func (r *Userrepository) DeleteUser(username string) error {
	query := `
	UPDATE users
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE username = $1
	`

	_, err := r.DB.Exec(query, username)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return err
	}

	return nil
}
