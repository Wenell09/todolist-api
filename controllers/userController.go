package controllers

import (
	"fmt"

	"github.com/Wenell09/todolist-api/database"
	"github.com/Wenell09/todolist-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseDataUserId struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	UserId  string `json:"user_id"`
}

type ResponseData struct {
	Status string      `json:"status"`
	Data   models.User `json:"data"`
}

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	id := uuid.New().String()
	results := new(models.ReqUser)
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
	registUser := models.User{
		UserId:   id,
		Username: results.Username,
		Email:    results.Email,
		Password: string(hash),
	}
	if err := database.DB.Where("email = ?", registUser.Email).First(&user).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(ResponseMessage{
			Status:  "error",
			Message: "Email has been registered!",
		})
	}

	if err := database.DB.Model(&user).Create(&registUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
			Status:  "error",
			Message: "Error register user!",
		})
	}
	return c.JSON(ResponseDataUserId{Status: "success", Message: "Success register user!", UserId: id})
}

func LoginUser(c *fiber.Ctx) error {
	var user models.User
	results := new(models.ReqUser)
	if err := c.BodyParser(&results); err != nil {
		return err
	}
	if results.Email == "" || results.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseMessage{
			Status:  "error",
			Message: "make sure all body fields are filled in!",
		})
	}
	loginUser := models.User{
		Email:    results.Email,
		Password: results.Password,
	}

	if err := database.DB.Where("email = ?", loginUser.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
			Status:  "error",
			Message: "make sure your email is registered!",
		})
	}
	if matchPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); matchPass != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
			Status:  "error",
			Message: "make sure your email and password correct!",
		})
	}
	return c.JSON(ResponseDataUserId{
		Status:  "success",
		Message: fmt.Sprintf("Welcome %s", user.Username),
		UserId:  user.UserId,
	})
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	userId := c.Params("user_id")
	if err := database.DB.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
			Status:  "error",
			Message: fmt.Sprintf("User with id:%s not found", userId),
		})
	}
	return c.JSON(ResponseData{
		Status: "success",
		Data:   user,
	})
}

func EditUser(c *fiber.Ctx) error {
	var user models.User
	results := new(models.ReqUser)
	userId := c.Params("user_id")
	if err := c.BodyParser(&results); err != nil {
		return err
	}
	if results.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(results.Password), bcrypt.DefaultCost)
		editUser := models.User{
			Username: results.Username,
			Email:    results.Email,
			Password: string(hash),
		}
		if err := database.DB.Where("user_id = ?", userId).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
				Status:  "error",
				Message: fmt.Sprintf("User with id:%s not found,failed to update!", userId),
			})
		}
		if err := database.DB.Model(&user).Where("user_id = ?", userId).Updates(&editUser).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ResponseMessage{
				Status:  "error",
				Message: fmt.Sprintf("user with id:%s update failed!", userId),
			})
		}
	} else {
		editUser := models.User{
			Username: results.Username,
			Email:    results.Email,
		}
		if err := database.DB.Where("user_id = ?", userId).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
				Status:  "error",
				Message: fmt.Sprintf("User with id:%s not found,failed to update!", userId),
			})
		}
		if err := database.DB.Model(&user).Where("user_id = ?", userId).Updates(&editUser).Error; err != nil {
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
	var user models.User
	userId := c.Params("user_id")
	if err := database.DB.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{
			Status:  "error",
			Message: fmt.Sprintf("user with id:%s not found,failed to delete!", userId),
		})
	}
	if err := database.DB.Model(&user).Where("user_id = ?", userId).Delete(&user).Error; err != nil {
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
