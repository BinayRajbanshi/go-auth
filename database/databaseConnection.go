package database

import (
	"context"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error laoding .env file")
	}

	mongoUrl := os.Getenv("MONGODB_URL")

	ctx, cancelCtx := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelCtx()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Monogo db connected successfully.")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster-1").Collection(collectionName)
	return collection
}
