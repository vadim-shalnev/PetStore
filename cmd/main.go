package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vadim-shalnev/PetStore/config"
	"github.com/vadim-shalnev/PetStore/internal/Pet/petController"
	"github.com/vadim-shalnev/PetStore/internal/Pet/petRepository"
	"github.com/vadim-shalnev/PetStore/internal/Pet/petService"
	"github.com/vadim-shalnev/PetStore/internal/Store/storeController"
	"github.com/vadim-shalnev/PetStore/internal/Store/storeRepository"
	"github.com/vadim-shalnev/PetStore/internal/Store/storeService"
	"github.com/vadim-shalnev/PetStore/internal/User/userController"
	"github.com/vadim-shalnev/PetStore/internal/User/userRepository"
	"github.com/vadim-shalnev/PetStore/internal/User/userService"
	"github.com/vadim-shalnev/PetStore/internal/handler"
	"github.com/vadim-shalnev/PetStore/models/controllers"
	"log"
	"net/http"
	"time"
)

// @title PetStore API
// @version 1.0
// @description This is a geocode api server.

// @host localhost:8080
// @BasePath /api/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	// Создаем конфигурацию приложения
	conf := config.NewAppConf()

	db := ConnectionDB(conf)
	defer db.Close()
	// Создаем таблицу
	CreateTable(db)
	CreateCategories(db)

	// Инициализируем слои и запускаем хэндлер
	st := ProjectInit(db)
	router := handler.InitRouters(st)
	http.ListenAndServe(":8080", router)
}

// ConnectionDB Подключаемся к бд
func ConnectionDB(conf config.AppConf) *sql.DB {
	time.Sleep(time.Second * 2)
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

func ProjectInit(db *sql.DB) controllers.Controllers {
	userrepository := userRepository.NewUserRepository(db)
	storerepository := storeRepository.NewStoreRepository(db)
	petrepository := petRepository.NewPetRepository(db)
	userservice := userService.NewUserService(userrepository)
	storesrervice := storeService.NewStoreService(storerepository)
	petserervice := petService.NewPetService(petrepository)
	userrontroller := userController.NewUserController(userservice)
	storerontroller := storeController.NewStoreController(storesrervice)
	petrontroller := petController.NewPetController(petserervice)
	return controllers.Controllers{
		User:  userrontroller,
		Store: storerontroller,
		Pet:   petrontroller,
	}

}
