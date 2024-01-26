package routes

import (
	"github.com/goravel/framework/contracts/http"

	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Post("/users", userController.Store)
	facades.Route().Get("/users/{id}", userController.Show)
	facades.Route().Delete("/users/{id}", userController.Delete)

	facades.Route().Get("/health", func(ctx http.Context) http.Response {
		return ctx.Response().Json(200, http.Json{
			"Hello": "Goravel",
		})
	})
}
