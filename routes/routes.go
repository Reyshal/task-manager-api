package routes

import (
	"github.com/Reyshal/task-manager-api/config"
	"github.com/Reyshal/task-manager-api/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "So far so good 🚀",
		})
	})

	// Auth routes
	auth := app.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Protected routes
	api := app.Group("/api", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(config.ConfigInstance.JWT.SecretKey),
		},
	}))

	// Task routes
	api.Get("/tasks", handlers.GetTasks)
	api.Post("/tasks", handlers.CreateTask)
	api.Put("/tasks/:id", handlers.UpdateTask)
	api.Delete("/tasks/:id", handlers.DeleteTask)
}
