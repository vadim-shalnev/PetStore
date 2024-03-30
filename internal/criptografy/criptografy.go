package criptografy

import (
	"github.com/vadim-shalnev/PetStore/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(user *models.User) error {
	password := user.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
func CheckPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("Неверный пароль", hashedPassword, password)
		return err
	}
	return nil
}
