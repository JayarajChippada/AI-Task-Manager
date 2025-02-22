package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"backend/config"
	"backend/models"
)

// CreateTask - Creates a new task
func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	task.ID = primitive.NewObjectID()
	task.Status = models.Pending
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := config.TasksCollection.InsertOne(ctx, task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
	}

	return c.Status(http.StatusCreated).JSON(task)
}

// GetAllTasks - Fetches all tasks
func GetAllTasks(c *fiber.Ctx) error {
	var tasks []models.Task

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.TasksCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch tasks"})
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error decoding task"})
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

// GetMyTasks - Fetches tasks assigned to the logged-in user
func GetMyTasks(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string) // Extract user ID from middleware

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var tasks []models.Task

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.TasksCollection.Find(ctx, bson.M{"assigned_to": objID})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch tasks"})
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error decoding task"})
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

// GetTaskByID - Fetch a specific task by ID
func GetTaskByID(c *fiber.Ctx) error {
	taskID := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	var task models.Task

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = config.TasksCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	} else if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch task"})
	}

	return c.JSON(task)
}

// UpdateTask - Update task details
func UpdateTask(c *fiber.Ctx) error {
	taskID := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	var updateData models.Task
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request data"})
	}

	updateData.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": updateData,
	}

	_, err = config.TasksCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update task"})
	}

	return c.JSON(fiber.Map{"message": "Task updated successfully"})
}

// UpdateTaskStatus - Updates only the task status
func UpdateTaskStatus(c *fiber.Ctx) error {
	taskID := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	var statusUpdate struct {
		Status models.TaskStatus `json:"status"`
	}
	if err := c.BodyParser(&statusUpdate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request data"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{"status": statusUpdate.Status, "updated_at": time.Now()},
	}

	_, err = config.TasksCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update task status"})
	}

	return c.JSON(fiber.Map{"message": "Task status updated successfully"})
}

// DeleteTask - Deletes a task by ID
func DeleteTask(c *fiber.Ctx) error {
	taskID := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = config.TasksCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete task"})
	}

	return c.JSON(fiber.Map{"message": "Task deleted successfully"})
}
