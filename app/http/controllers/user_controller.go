package controllers

import (
	"fmt"
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
	validator, err := ctx.Request().Validate(map[string]string{
		"name":     "required|max_len:255",
		"username": "required|max_len:255",
		"password": "required|max_len:255",
	})

	if validator.Fails() {
		return ctx.Response().Success().Json(http.Json{
			"message": validator.Errors().All(),
		})
	}

	fmt.Printf(err.Error())

	name := ctx.Request().Input("name")
	username := ctx.Request().Input("username")
	password := ctx.Request().Input("password")

	user := models.User{Name: name, Username: username, Password: password}
	result := facades.Orm().Query().Create(&user)

	return ctx.Response().Success().Json(http.Json{
		"message": result,
	})
}

func (r *UserController) Show(ctx http.Context) http.Response {
	var user models.User
	id := ctx.Request().Input("id")

	facades.Orm().Query().FindOrFail(&user, id)

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
