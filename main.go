package main

import (
	"github.com/Wenell09/todolist-api/database"
	"github.com/Wenell09/todolist-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	routes.RouteApp(app)
	app.Listen(":8000")
}
