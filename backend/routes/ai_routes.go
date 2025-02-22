package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/controllers"
)

func SetupAIRoutes(app *fiber.App) {
	ai := app.Group("/ai")

	ai.Post("/suggest-tasks", controllers.GenerateTaskSuggestions)
	ai.Post("/improve-task", controllers.ImproveTask)
	ai.Post("/assign-priority", controllers.AssignTaskPriority)
}
