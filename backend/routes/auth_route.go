package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/controllers"
)

func SetupAuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}
