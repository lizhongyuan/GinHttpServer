package common


import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
	"log"
)


var DB *mongo.Database


func InitMongoDatabase() *mongo.Database {
	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
	if err != nil {
		log.Fatal(err)
	}

	DB := client.Database("city_mocha")

	return DB
}

