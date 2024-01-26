package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Store(ctx http.Context) http.Response {
	validator, _ := ctx.Request().Validate(map[string]string{
		"name":     "required|max_len:255",
		"username": "required|max_len:255",
		"password": "required|max_len:255",
	})

	if validator.Fails() {
		return ctx.Response().Success().Json(http.Json{
			"message": validator.Errors().All(),
		})
	}

	name := ctx.Request().Input("name")
	username := ctx.Request().Input("username")
	password := ctx.Request().Input("password")

	user := models.User{Name: name, Username: username, Password: password}
	facades.Orm().Query().Where("username", username).FirstOrCreate(&user, user)

	return ctx.Response().Status(201).Json(http.Json{
		"message": &user,
	})
}

func (r *UserController) Show(ctx http.Context) http.Response {
	var user models.User
	id := ctx.Request().Input("id")

	err := facades.Orm().Query().FindOrFail(&user, id)

	if err != nil {
		return ctx.Response().Status(404).Json(http.Json{
			"message": "Id não encontrado na base de dados",
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"data": &user,
	})
}

func (r *UserController) Delete(ctx http.Context) http.Response {
	var user models.User
	id := ctx.Request().Input("id")

	facades.Orm().Query().Delete(&user, id)

	return ctx.Response().Status(204).Json(http.Json{})
}
