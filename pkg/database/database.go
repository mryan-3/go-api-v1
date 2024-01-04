package database

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

var collection *mongo.Collection
var ctx = context.TODO()

func ConnectDB() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
         log.Fatal("Error loading .env file")
    }

    // Get URI
    uri := os.Getenv("MONGODB_URI")
    if uri == "" {
         log.Fatal("MONGODB_URI must be set")
    }

    // Connect with context timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    fmt.Println("Succesful connection")

    if err != nil {
         log.Fatal(err)
    }

    // Disconnect with timeout
    ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := client.Disconnect(ctx); err != nil {
     log.Fatal(err)
    }
}
