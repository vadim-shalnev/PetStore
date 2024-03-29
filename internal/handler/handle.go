package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/vadim-shalnev/PetStore/internal/Store/storeController"
	"github.com/vadim-shalnev/PetStore/internal/User/userController"
	"github.com/vadim-shalnev/PetStore/internal/middleware"
	"net/http"
)

// маршрутизатор принимает структуры слоев

func InitRouters(userController *userController.UserController, storeController *storeController.StoreController, petController *etController.PetController) http.Handler {
	r := chi.NewRouter()
	// User
	controller := userController
	r.Route("/api/", func(r chi.Router) {
		//r.Use(controller.AuthMiddleware)
		r.Post("/user", controller.CreateUser)
		r.Post("/user/createWithList", controller.CreateUsers)
		r.Get("/user/login", controller.Login)
		r.Get("/user/logout", controller.Logout)
		r.Get("/user/{username}", controller.GetUser)
		r.Put("/user/{username}", controller.UpdateUser)
		r.Delete("/user/{username}", controller.DeleteUser)
	})
	// Store
	controller = storeController
	r.Route("/api/store/", func(r chi.Router) {
		r.Post("/order", controller.NewOrder)
		r.Get("/order/{orderId}", controller.GetOrder)
		r.Delete("/order/{orderId}", controller.DeleteOrder)
		r.Use(middleware.RefreshToken)
		r.Get("/order/inventory", controller.Getinventory)
	})
	// Pets
	controller = petController
	r.Route("/api/", func(r chi.Router) {
		r.Use(middleware.RefreshToken)
		//r.Post("/pet/{petId}/uploadImage",controller.UpImage)
		r.Post("/pet", controller.AddPet)
		r.Put("/pet", controller.UpdatePet)
		r.Get("/pet/findByStatus", controller.FindByStatus)
		r.Get("/pet/{petId}", controller.GetPet)
		r.Post("/pet/{petId}", controller.ChangePet)
		r.Delete("/pet/{petId}", controller.DeletePet)
	})
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return r
}
