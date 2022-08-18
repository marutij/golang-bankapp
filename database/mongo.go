package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context = context.Background()
var uri = "mongodb://localhost:27017"
var database = "Bank"
var client *mongo.Client

func init() {
	client = NewDbClient()
}

func NewDbClient() *mongo.Client {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Error while creating mongo client")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("Error while connecting database", err.Error())
		client.Disconnect(ctx)
	}

	log.Println("Database Connected")
	return client
}

func GetCollectionInstance(collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}
