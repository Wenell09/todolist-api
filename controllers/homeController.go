package controllers

import "github.com/gofiber/fiber/v2"

func HomeController(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Welcome to todolist-api",
	})
}
