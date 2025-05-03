package controllers

import (
	"fmt"

	"github.com/Wenell09/todolist-api/database"
	"github.com/Wenell09/todolist-api/models/user"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ResponseDataUserId struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	UserId  string `json:"user_id"`
}

type ResponseData struct {
	Status string    `json:"status"`
	Data   user.User `json:"data"`
}

func RegisterUser(c *fiber.Ctx) error {
	var data user.User
	id := uuid.New().String()
	results := new(user.ReqUser)
	if err := c.BodyParser(&results); err != nil {
		return err
	}
	if results.Username == "" || results.Email == "" || results.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseMessage{
			Status:  "error",
			Message: "make sure all body fields are filled in!",
		})
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(results.Password), bcrypt.DefaultCost)
	registUser := user.User{
		UserId:   id,
		Username: results.Username,
		Email:    results.Email,
		Password: string(hash),
	}
	if err := database.DB.Where("email = ?", registUser.Email).First(&data).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(ResponseMessage{
			Status:  "error",
			Message: "Email has been registered!",
		})
	}

	if err := database.DB.Model(&data).Create(&registUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
			Status:  "error",
			Message: "Error register user!",
		})
	}
	return c.JSON(ResponseDataUserId{Status: "success", Message: "Success register user!", UserId: id})
}

func LoginUser(c *fiber.Ctx) error {
	var data user.User
	results := new(user.ReqUser)
	if err := c.BodyParser(&results); err != nil {
		return err
	}
	if results.Email == "" || results.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseMessage{
			Status:  "error",
			Message: "make sure all body fields are filled in!",
		})
	}
	loginUser := user.User{
		Email:    results.Email,
		Password: results.Password,
	}

	if err := database.DB.Where("email = ?", loginUser.Email).First(&data).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
			Status:  "error",
			Message: "make sure your email is registered!",
		})
	}
	if matchPass := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(loginUser.Password)); matchPass != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
			Status:  "error",
			Message: "make sure your email and password correct!",
		})
	}
	return c.JSON(ResponseDataUserId{
		Status:  "success",
		Message: fmt.Sprintf("Welcome %s", data.Username),
		UserId:  data.UserId,
	})
}

func GetUser(c *fiber.Ctx) error {
	var data user.User
	userId := c.Params("user_id")
	if err := database.DB.Where("user_id = ?", userId).First(&data).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
			Status:  "error",
			Message: fmt.Sprintf("User with id:%s not found", userId),
		})
	}
	return c.JSON(ResponseData{
		Status: "success",
		Data:   data,
	})
}

func EditUser(c *fiber.Ctx) error {
	var data user.User
	results := new(user.ReqUser)
	userId := c.Params("user_id")
	if err := c.BodyParser(&results); err != nil {
		return err
	}
	if results.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(results.Password), bcrypt.DefaultCost)
		editUser := user.User{
			Username: results.Username,
			Email:    results.Email,
			Password: string(hash),
		}
		if err := database.DB.Where("user_id = ?", userId).First(&data).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
				Status:  "error",
				Message: fmt.Sprintf("User with id:%s not found,failed to update!", userId),
			})
		}
		if err := database.DB.Model(&data).Where("user_id = ?", userId).Updates(&editUser).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
				Status:  "error",
				Message: fmt.Sprintf("user with id:%s update failed!", userId),
			})
		}
	} else {
		editUser := user.User{
			Username: results.Username,
			Email:    results.Email,
		}
		if err := database.DB.Where("user_id = ?", userId).First(&data).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
				Status:  "error",
				Message: fmt.Sprintf("User with id:%s not found,failed to update!", userId),
			})
		}
		if err := database.DB.Model(&data).Where("user_id = ?", userId).Updates(&editUser).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
				Status:  "error",
				Message: fmt.Sprintf("user with id:%s failed to update!", userId),
			})
		}
	}
	return c.JSON(ResponseMessage{
		Status:  "success",
		Message: fmt.Sprintf("user with id:%s successfully updated!", userId),
	})
}

func DeleteUser(c *fiber.Ctx) error {
	var data user.User
	userId := c.Params("user_id")
	if err := database.DB.Where("user_id = ?", userId).First(&data).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
			Status:  "error",
			Message: fmt.Sprintf("user with id:%s not found,failed to delete!", userId),
		})
	}
	if err := database.DB.Model(&data).Where("user_id = ?", userId).Delete(&data).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
			Status:  "error",
			Message: fmt.Sprintf("user with id:%s failed to delete!", userId),
		})
	}
	return c.JSON(ResponseMessage{
		Status:  "success",
		Message: fmt.Sprintf("user with id:%s successfully deleted", userId),
	})
}
