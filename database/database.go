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

func ConnectDB() (*mongo.Client, error) {
	_ = godotenv.Load()
	MongoURL := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoURL))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}
