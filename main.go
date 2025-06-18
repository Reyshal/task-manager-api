package main

import (
	"log"

	"github.com/Reyshal/task-manager-api/config"
	"github.com/Reyshal/task-manager-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	log.Println("🚀 Starting the server...")

	// Load config and database
	config.InitConfig()
	config.InitDatabase()

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Authorization",
	}))

	// Setup routes
	routes.SetupRoutes(app)

	// Health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "✅ Task Manager API is running!",
		})
	})

	// Start the server
	log.Println("✅ Server started on port " + config.ConfigInstance.Server.Port)
	app.Listen(":" + config.ConfigInstance.Server.Port)
}
