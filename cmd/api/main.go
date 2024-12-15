// package main

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/1kazuchi/go-mongo-todos/db"
// 	"github.com/1kazuchi/go-mongo-todos/handlers"
// 	"github.com/1kazuchi/go-mongo-todos/services"
// )

// type Application struct {
// 	Models services.Models
// }

// func main() {
// 	mongoClient, err := db.ConnectToMongo()
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	defer func() {
// 		if err = mongoClient.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	services.New(mongoClient)

// 	log.Println("Server running in port", 8080)
// 	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))

// }

package main

import (
	"log"
	"net/http"

	"github.com/1kazuchi/go-mongo-todos/db"
	"github.com/1kazuchi/go-mongo-todos/handlers"
	"github.com/1kazuchi/go-mongo-todos/services"
)

func main() {
	// Call ConnectToMongo from the db package
	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Panic("Failed to connect to MongoDB:", err)
	}

	// Initialize your services with the mongoClient
	services.New(mongoClient)

	// Setting up the HTTP server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))
}
