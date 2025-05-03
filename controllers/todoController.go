package controllers

import (
	"fmt"

	"github.com/Wenell09/todolist-api/database"
	"github.com/Wenell09/todolist-api/models/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ResponseDataTodo struct {
	Status string      `json:"status"`
	Data   []todo.Todo `json:"data"`
}

func AddTodo(c *fiber.Ctx) error {
	var data todo.Todo
	todoId := uuid.New().String()
	results := new(todo.ReqTodo)
	if err := c.BodyParser(&results); err != nil {
		return err
	}
	if results.TodoTitle == "" || results.TodoDesc == "" || results.UserId == "" || results.PrioritasId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseMessage{
			Status:  "error",
			Message: "make sure all body fields are filled in!",
		})
	}
	addTodo := todo.Todo{
		TodoId:      todoId,
		TodoTitle:   results.TodoTitle,
		TodoDesc:    results.TodoDesc,
		UserId:      results.UserId,
		PrioritasId: results.PrioritasId,
	}
	if err := database.DB.Model(&data).Create(&addTodo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
			Status:  "error",
			Message: "failed to create a new todo!",
		})
	}
	return c.JSON(ResponseMessage{
		Status:  "success",
		Message: "successfully created a new todo",
	})
}

func GetTodo(c *fiber.Ctx) error {
	var data []todo.Todo
	userId := c.Params("user_id")

	if err := database.DB.Preload("Prioritas").Where("user_id = ?", userId).Find(&data).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
			Status:  "error",
			Message: fmt.Sprintf("todo owned by user id:%s not found", userId),
		})
	}

	return c.JSON(ResponseDataTodo{
		Status: "success",
		Data:   data,
	})
}
