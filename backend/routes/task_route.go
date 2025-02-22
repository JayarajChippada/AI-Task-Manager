package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/controllers"
	"backend/middleware"
)

func SetupTaskRoutes(app *fiber.App) {
	task := app.Group("/tasks", middleware.AuthMiddleware)

	task.Post("/", controllers.CreateTask)          // Create a new task
	task.Get("/", controllers.GetAllTasks)         // Get all tasks
	task.Get("/assigned", controllers.GetMyTasks)  // Get tasks assigned to the logged-in user
	task.Get("/:id", controllers.GetTaskByID)      // Get task by ID
	task.Put("/:id", controllers.UpdateTask)       // Update task details
	task.Patch("/:id/status", controllers.UpdateTaskStatus) // Update task status
	task.Delete("/:id", controllers.DeleteTask)    // Delete task
}
