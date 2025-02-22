package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var UsersCollection *mongo.Collection
var TasksCollection *mongo.Collection // Add this

func ConnectDB() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI not set in .env")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	DB = client
	UsersCollection = client.Database("taskapp").Collection("users")
	TasksCollection = client.Database("taskapp").Collection("tasks") // Add this

	log.Println("Connected to MongoDB")
}
