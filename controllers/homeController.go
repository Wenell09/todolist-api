package controllers

import "github.com/gofiber/fiber/v2"

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HomeController(c *fiber.Ctx) error {
	return c.JSON(ResponseMessage{
		Status:  "success",
		Message: "Welcome to todolist-api",
	})
}
