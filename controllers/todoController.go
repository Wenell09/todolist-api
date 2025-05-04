package controllers

import (
	"fmt"

	"github.com/Wenell09/todolist-api/database"
	"github.com/Wenell09/todolist-api/models/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ResponseDataTodo struct {
	Status string         `json:"status"`
	Data   []todo.ResTodo `json:"data"`
}

func AddTodo(c *fiber.Ctx) error {
	var model todo.Todo
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
	if err := database.DB.Model(&model).Create(&addTodo).Error; err != nil {
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
	var model todo.Todo
	var data []todo.ResTodo
	userId := c.Params("user_id")
	if err := database.DB.
		Model(&model).
		Select("todos.*, users.username, prioritas.prioritas_id, prioritas.prioritas_name").
		Joins("LEFT JOIN users ON users.user_id = todos.user_id").
		Joins("LEFT JOIN prioritas ON prioritas.prioritas_id = todos.prioritas_id").
		Where("todos.user_id = ?", userId).
		Scan(&data).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
			Status:  "error",
			Message: fmt.Sprintf("Todo with user id:%s not found", userId),
		})
	}
	return c.JSON(ResponseDataTodo{
		Status: "success",
		Data:   data,
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	var model todo.Todo
	results := new(todo.ReqTodo)
	todoId := c.Params("todo_id")
	if err := c.BodyParser(&results); err != nil {
		return err
	}
	updateTodo := todo.Todo{
		TodoTitle:   results.TodoTitle,
		TodoDesc:    results.TodoDesc,
		PrioritasId: results.PrioritasId,
	}
	if err := database.DB.First(&model, "user_id = ? AND todo_id = ?", results.UserId, todoId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
			Status:  "error",
			Message: fmt.Sprintf("todo with id:%s not found,failed to update!", todoId),
		})
	}

	if err := database.DB.Model(&model).Where("user_id = ? AND todo_id = ?", results.UserId, todoId).Updates(&updateTodo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
			Status:  "error",
			Message: fmt.Sprintf("todo with id:%s failed to update!", todoId),
		})
	}
	return c.JSON(ResponseMessage{
		Status:  "success",
		Message: fmt.Sprintf("todo with id:%s successfully updated", todoId),
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	var model todo.Todo
	userId := c.Params("user_id")
	todoId := c.Params("todo_id")

	if todoId != "" {
		if err := database.DB.First(&model, "user_id = ? AND todo_id = ?", userId, todoId).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
				Status:  "error",
				Message: fmt.Sprintf("todo with user_id:%s and todo_id:%s not found,failed to delete!", userId, todoId),
			})
		}
		if err := database.DB.Model(&model).Where("user_id = ? AND todo_id = ?", userId, todoId).Delete(&model).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
				Status:  "error",
				Message: fmt.Sprintf("todo with user_id:%s and todo_id:%s failed to delete!", userId, todoId),
			})
		}
		return c.JSON(ResponseMessage{
			Status:  "success",
			Message: fmt.Sprintf("todo with user_id:%s and todo_id:%s successfully deleted", userId, todoId),
		})
	} else {
		if err := database.DB.First(&model, "user_id = ?", userId).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
				Status:  "error",
				Message: fmt.Sprintf("todo with user_id:%s not found,failed to delete all!", userId),
			})
		}
		if err := database.DB.Model(&model).Where("user_id = ?", userId).Delete(&model).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
				Status:  "error",
				Message: fmt.Sprintf("todo with user_id:%s failed to delete all!", userId),
			})
		}
		return c.JSON(ResponseMessage{
			Status:  "success",
			Message: fmt.Sprintf("todo with user_id:%s successfully deleted all", userId),
		})
	}
}
