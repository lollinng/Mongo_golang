package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() *mongo.Client {

	ctx, _ := context.WithTimeout(context.Background(),
		10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(EnvMongoUri()))
	if err != nil {
		panic(err)
	}

	// ping the db
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDb")

	return client
}

// Client instance
var DB *mongo.Client = ConnectDb()

// get Database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("TRY").Collection(collectionName)

}
