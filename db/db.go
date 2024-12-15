// package db

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var collection *mongo.Collection

// func ConnectToMongo() (*mongo.Client, error) {
// 	// Getting Mongo URI from environment variable
// 	MONGO_URI := os.Getenv("MONGO_URI")
// 	if MONGO_URI == "" {
// 		log.Fatal("MONGO_URI is not set in environment variables.")
// 		return nil, fmt.Errorf("MONGO_URI is not set")
// 	}

// 	// MongoDB connection string
// 	clientOptions := options.Client().ApplyURI(MONGO_URI)

// 	// Creating a context with timeout (5 seconds in this case)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}

// 	// Ping the MongoDB server to check if the connection is successful
// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal("Unable to ping MongoDB: ", err)
// 		return nil, err
// 	}

// 	log.Println("Connected to MongoDB!")

//		return client, nil
//	}
//
// db/db.go
package db

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

// ConnectToMongo establishes a connection to MongoDB and returns the client
func ConnectToMongo() (*mongo.Client, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Get the MongoDB URI from the environment variable
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return nil, fmt.Errorf("MONGO_URI is not set in the environment variables")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Establish connection
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MongoDB")
	return client, nil
}
