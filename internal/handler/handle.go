package handler

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/vadim-shalnev/PetStore/internal/middleware"
	"github.com/vadim-shalnev/PetStore/models/controllers"
	"net/http"
)

func InitRouters(controllers controllers.Controllers) http.Handler {
	r := chi.NewRouter()

	r.Route("/api/", func(r chi.Router) {
		// User
		userController := controllers.User
		r.Route("/user/", func(r chi.Router) {
			r.Post("/user", userController.CreateUser)
			r.Post("/user/createWithList", userController.CreateUsers)
			r.Get("/user/login", userController.Login)
			r.Get("/user/logout", userController.Logout)
			r.Get("/user/{username}", userController.GetUser)
			r.Put("/user/{username}", userController.UpdateUser)
			r.Delete("/user/{username}", userController.DeleteUser)
		})
		// Store
		storeController := controllers.Store
		r.Route("/store/", func(r chi.Router) {
			r.Post("/order", storeController.NewOrder)
			r.Get("/order/{orderId}", storeController.GetOrder)
			r.Delete("/order/{orderId}", storeController.DeleteOrder)
		})
		r.Route("/api/store/inventory", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Get("/", storeController.Getinventory)
		})
		// Pets
		petsController := controllers.Pet
		r.Route("/pet/", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			//r.Post("/pet/{petId}/uploadImage",userController.UpImage)
			r.Post("/pet", petsController.AddPet)
			r.Put("/pet", petsController.UpdatePet)
			r.Get("/pet/findByStatus", petsController.FindByStatus)
			r.Get("/pet/{petId}", petsController.GetPet)
			r.Post("/pet/{petId}", petsController.ChangePet)
			r.Delete("/pet/{petId}", petsController.DeletePet)
		})
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		))

	})

	return r
}
