package main

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"log"
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Size defines the item size
type Size struct {
	H   int
	W   float64
	Uom string
}

// Item defines an item
type Item struct {
	OID  objectid.ObjectID	`bson:"_id,omitempty"` // omitempty not working
	Item string				`bson:"item"`

	/*
	Qty  int
	Tags []string
	Size Size
	*/
}

func main() {
	// connect to MongoDB
	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("gotest")
	inventory := db.Collection("inventory")

	/*
	// write document
	itemWrite := Item{Item: "canvas", Qty: 100, Tags: []string{"cotton"}, Size: Size{H: 28, W: 35.5, Uom: "cm"}}
	itemWrite.OID = objectid.New()
	fmt.Printf("itemWrite = %v\n", itemWrite)

	result, err := inventory.InsertOne(context.Background(), itemWrite)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result = %#v\n", result)
	*/

	// read documents
	cursor, err := inventory.Find(
		context.Background(),
		bson.NewDocument(bson.EC.String("item", "canvas")),
	)
	if err != nil {
		log.Fatal(err)
	}
	// defer cursor.Close(context.Background())

	itemRead := Item{}
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&itemRead)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("itemRead = %v\n", itemRead)
	}
}
