package model

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
	"log"
)

type manager struct {
	name		string			`bson:"name"`
	contacts	string			`bson:"contacts"`
}

type User struct {
	OID				objectid.ObjectID		`json:"id" bson:"_id"`
	Roles			string					`bson:"roles"`
	Manager			manager					`bson:"manager"`
}


// todo: unfinish
func GetUserModel() {
	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
	if err != nil {
	log.Fatal(err)
	}
	db := client.Database("city_mocha")
	user := db.Collection("users")

	return user
}
