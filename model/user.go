package model


import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"../common"
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



func GetUserCollection() *mongo.Collection {

	// db, _ = common.GetMongoDatabase()

	collection := common.DB.Collection("users")

	return collection
}
