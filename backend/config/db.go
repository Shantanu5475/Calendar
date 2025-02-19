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

func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return fmt.Errorf("MONGO_URI is not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var connErr error
	client, connErr = mongo.Connect(ctx, clientOptions)
	if connErr != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", connErr)
	}

	pingErr := client.Ping(ctx, nil)
	if pingErr != nil {
		return fmt.Errorf("MongoDB ping failed: %w", pingErr)
	}

	fmt.Println("🚀 Connected to MongoDB successfully!")
	return nil
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
