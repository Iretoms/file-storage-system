package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

func Connect() *mongo.Client {
	clientOnce.Do(func() {
		fmt.Println("Initializing MongoDB connection")

		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		mongodbUri := os.Getenv("MONGODB_URI")

		if mongodbUri == "" {
			log.Fatal("MONGODB_URI environment variable is not set")
		}

		ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelCtx()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatalf("Failed to ping MongoDB: %v", err)
		}

		fmt.Println("Connected to MongoDB")
		clientInstance = client
	})

	return clientInstance
}

var DB *mongo.Client = Connect()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("file-storage-system").Collection(collectionName)
	return collection
}
