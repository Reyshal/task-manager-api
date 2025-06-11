package main

import (
	"log"

	"github.com/Reyshal/task-manager-api/config"
	"github.com/Reyshal/task-manager-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("🚀 Starting the server...")

	// Load config
	config.InitConfig()

	// Connect to the database
	database.InitDatabase()

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

	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	// Start the server
	log.Println("✅ Server started on port " + config.ConfigInstance.Server.Port)
	app.Listen(":" + config.ConfigInstance.Server.Port)
}
