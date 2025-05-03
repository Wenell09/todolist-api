package routes

import (
	"github.com/Wenell09/todolist-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteApp(app *fiber.App) {
	app.Get("/", controllers.HomeController)
	app.Post("/api/register", controllers.RegisterUser)
	app.Post("/api/login", controllers.LoginUser)
	app.Get("/api/user/:user_id", controllers.GetUser)
	app.Patch("/api/editUser/:user_id", controllers.EditUser)
	app.Delete("/api/deleteUser/:user_id", controllers.DeleteUser)
}
