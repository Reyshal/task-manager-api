package routes

import (
	"github.com/Reyshal/task-manager-api/config"
	"github.com/Reyshal/task-manager-api/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Auth routes
	auth := app.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// TODO: create Protected task routes
	api := app.Group("/api", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(config.ConfigInstance.JWT.SecretKey),
		},
	}))
	api.Get("/authorized", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Authorized",
		})
	})
}
