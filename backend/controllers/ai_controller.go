package controllers

import (
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

// GenerateTaskSuggestions - Get AI-generated tasks for a project
func GenerateTaskSuggestions(c *fiber.Ctx) error {
	var request struct {
		ProjectDescription string `json:"project_description"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	suggestions, err := services.GetTaskSuggestions(request.ProjectDescription)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate suggestions"})
	}

	return c.JSON(fiber.Map{"suggestions": suggestions})
}

// ImproveTask - Enhance task descriptions using AI
func ImproveTask(c *fiber.Ctx) error {
	var request struct {
		TaskDescription string `json:"task_description"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	improvedTask, err := services.ImproveTaskDescription(request.TaskDescription)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to improve task"})
	}

	return c.JSON(fiber.Map{"improved_task": improvedTask})
}

// AssignTaskPriority - AI assigns priority levels to tasks
func AssignTaskPriority(c *fiber.Ctx) error {
	var request struct {
		TaskDescription string `json:"task_description"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	priority, err := services.AssignTaskPriority(request.TaskDescription)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to assign priority"})
	}

	return c.JSON(fiber.Map{"priority": priority})
}
