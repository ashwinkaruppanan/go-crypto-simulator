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

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}

	MongoURL := os.Getenv("MONGO")

	client, clientErr := mongo.NewClient(options.Client().ApplyURI(MongoURL))
	if clientErr != nil {
		log.Panic(clientErr)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	if connectionErr := client.Connect(ctx); connectionErr != nil {
		log.Fatal(connectionErr)
	}
	defer cancel()

	fmt.Println("DB connected successfully")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("go-crypto").Collection(collectionName)
}
