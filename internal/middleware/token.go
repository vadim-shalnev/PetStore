package middleware

import (
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vadim-shalnev/PetStore/models"
	"log"
	"net/http"
	"strings"
	"time"
)

func NewTokenMiddleware(user *models.User) (string, error) {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"ID":       user.ID,
		"Username": user.Username,
		"Exp":      time.Now().Add(time.Second * 60).Unix(),
	})
	if err != nil {
		return "", errors.New("token generation error")
	}
	return tokenString, nil
}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Usertoken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		token := VerifyToken(Usertoken)
		log.Println("midleware")
		if !token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func VerifyToken(tokenString string) bool {
	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Возвращаем секретный ключ для проверки подписи
		return []byte("secret"), nil
	})
	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Проверяем срок действия токена
		exp := int64(claims["Exp"].(float64))
		if time.Now().Unix() > exp {
			return false
		}
		return true
	}
	return false
}

func GetUserinfo(tokenstring string) (string, string) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		// Возвращаем секретный ключ для проверки подписи
		return []byte("secret"), nil
	})
	if err != nil {
		return "", ""
	}
	if !token.Valid {
		return "", ""
	}

	var username string
	var id string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		id = claims["ID"].(string)
		username = claims["Username"].(string)

		return id, username
	}
	return "", ""
}
