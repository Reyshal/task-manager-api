package handlers

import (
	"github.com/Reyshal/task-manager-api/dto"
	"github.com/Reyshal/task-manager-api/models"
	"github.com/Reyshal/task-manager-api/services"
	"github.com/Reyshal/task-manager-api/utils"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data dto.RegisterRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check if user already exists
	userService := services.NewUserService()
	if user, _ := userService.GetByEmail(data.Email); user != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Create user
	user := models.User{
		Username: data.Username,
		Email:    data.Email,
		Password: hashedPassword,
	}
	if err := userService.Create(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func Login(c *fiber.Ctx) error {
	var data dto.LoginRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check if user exists
	userService := services.NewUserService()
	user, err := userService.GetByEmail(data.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Check password
	if !utils.CheckPasswordHash(data.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Generate JWT token
	token, err := utils.GenerateToken(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
