package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongo.Database


func ConnectToMongo() {

	// Check the Environment  Variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Set Environment  Variable
	uri := os.Getenv("DB_URI")

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		 panic(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
	 panic(err)
	}

	DB = *client.Database("TareasDB") 

	fmt.Println("Connected to MongoDB")
}
