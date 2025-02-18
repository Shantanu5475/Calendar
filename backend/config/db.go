package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var connErr error
	client, connErr = mongo.Connect(ctx, clientOptions)
	if connErr != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", connErr)
	}

	pingErr := client.Ping(ctx, nil)
	if pingErr != nil {
		log.Fatalf("MongoDB ping failed: %v", pingErr)
	}

	fmt.Println("🚀 Connected to MongoDB successfully!")
}

func GetCollection(collectionName string) *mongo.Collection {
	if client == nil {
		log.Fatal("MongoDB client is not initialized. Did you call ConnectDB()?")
	}
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		log.Fatal("DATABASE_NAME is not set in .env file")
	}
	return client.Database(databaseName).Collection(collectionName)
}
