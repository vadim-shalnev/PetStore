package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vadim-shalnev/PetStore/config"
	"github.com/vadim-shalnev/PetStore/internal/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	// Создаем конфигурацию приложения
	conf := config.NewAppConf()

	db := ConnectionDB(conf)
	// Создаем таблицу
	CreateTable(db)
	// Инициализируем слои и запускаем хэндлер
	st := ProjectInit(db)
	router := handler.InitRouters(st)
	http.ListenAndServe(":8080", router)
}

// ConnectionDB Подключаемся к бд
func ConnectionDB(conf config.AppConf) *sql.DB {
	time.Sleep(time.Second * 5)
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
	return db
}

type Controllers struct {
	User
	Store
	Pet
}

func ProjectInit(db *sql.DB) Controllers {

}
