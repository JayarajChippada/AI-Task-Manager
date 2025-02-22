package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskStatus string

const (
	Pending    TaskStatus = "pending"
	InProgress TaskStatus = "in_progress"
	Completed  TaskStatus = "completed"
)

type PriorityLevel string

const (
	Low    PriorityLevel = "low"
	Medium PriorityLevel = "medium"
	High   PriorityLevel = "high"
	Urgent PriorityLevel = "urgent"
)

type Comment struct {
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Text      string             `bson:"text" json:"text"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type Task struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title       string               `bson:"title" json:"title" validate:"required,min=3,max=100"`
	Description string               `bson:"description,omitempty" json:"description,omitempty"`
	AssignedTo  []primitive.ObjectID `bson:"assigned_to,omitempty" json:"assigned_to,omitempty"` // Multiple users
	Status      TaskStatus           `bson:"status" json:"status" validate:"required,oneof=pending in_progress completed"`
	Priority    PriorityLevel        `bson:"priority" json:"priority" validate:"oneof=low medium high urgent"`
	DueDate     *time.Time           `bson:"due_date,omitempty" json:"due_date,omitempty"`
	Comments    []Comment            `bson:"comments,omitempty" json:"comments,omitempty"` // Task discussion
	CreatedAt   time.Time            `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt   time.Time            `bson:"updated_at,omitempty" json:"updated_at"`
}
