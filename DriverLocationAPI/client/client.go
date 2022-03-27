package client

import (
	"bitaksi/config"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client     *mongo.Client
	DB         *mongo.Database
	Collection *mongo.Collection
}

var BitaksiInstance *MongoInstance

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MONGO["URI"]))
	if err != nil {
		log.Fatal("Error while connection mongoDB, please check your URI", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB.")

	BitaksiInstance = &MongoInstance{
		Client: client,
		DB:     client.Database(config.MONGO["DATABASE"]),
	}
	BitaksiInstance.Collection = BitaksiInstance.DB.Collection(config.MONGO["COLLECTION"])
}
