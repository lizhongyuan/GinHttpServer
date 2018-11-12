package model


import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"../common"
)


type Manager struct {
	name		string			`json:"name" bson:"name"`
	contacts	string			`json:"contacts" bson:"contacts"`
}


type User struct {
	OID				objectid.ObjectID		`json:"id" bson:"_id"`
	Roles			string					`bson:"roles"`
	Manager			Manager					`bson:"manager"`
}



func GetUserCollection() *mongo.Collection {

	collection := common.DB.Collection("users")

	return collection
}
