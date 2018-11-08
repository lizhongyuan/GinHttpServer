package model

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
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
