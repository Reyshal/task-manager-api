package handlers

import (
	"github.com/Reyshal/task-manager-api/dto"
	"github.com/Reyshal/task-manager-api/models"
	"github.com/Reyshal/task-manager-api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetTasks(c *fiber.Ctx) error {
	// Get user id from jwt
	user := c.Locals("user").(*jwt.Token)
	userID := uint(user.Claims.(jwt.MapClaims)["user_id"].(float64))

	// Get tasks
	taskService := services.NewTaskService()
	tasks, err := taskService.GetTasks(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get tasks",
		})
	}

	return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	// Get body request
	var data dto.CreateTaskRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Get user id from jwt
	user := c.Locals("user").(*jwt.Token)
	userID := uint(user.Claims.(jwt.MapClaims)["user_id"].(float64))

	// Create task
	taskService := services.NewTaskService()
	task := models.Task{
		Title:  data.Title,
		UserID: userID,
	}

	if err := taskService.Create(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create task",
		})
	}

	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	// Get task id
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Parameter 'id' is required",
		})
	}

	// Get body request
	var data dto.UpdateTaskRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Find task
	taskService := services.NewTaskService()
	task, err := taskService.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	// Update task
	task.Title = data.Title
	if err := taskService.Update(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update task",
		})
	}

	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	// Get task id
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Parameter 'id' is required",
		})
	}

	// Find task
	taskService := services.NewTaskService()
	task, err := taskService.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	// Delete task
	if err := taskService.Delete(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete task",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}
