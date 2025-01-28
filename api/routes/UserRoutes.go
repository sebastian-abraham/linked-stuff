package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString("All users")
	})
}
