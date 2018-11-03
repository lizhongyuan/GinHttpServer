package main


import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"time"
)

type inspectDeviceSummaryItem struct {
	DeviceType	string
	Num			int
}

type unionDeviceSummaryItem struct {
	UnionType	string
	Num			int
}


/*
type inspectTask struct {
	OID	objectid.ObjectID	`bson:_id, omitempty`
	Name		string
	Identifier	string		`bson:identifier`
}
*/


type inspectTask struct {
	OID			objectid.ObjectID		`bson:"_id,omitempty"`
	Name		string				`bson:"name"`
	Identifier	string				`bson:"identifier"`

	Status			int					`bson:"status"`
	CreatedBy		objectid.ObjectID	`bson:"createdBy"`
	CreatedTime		time.Time			`bson:"createdTime"`
	BeginTime		time.Time			`bson:"beginTime"`
	EndTime			time.Time			`bson:"endTime"`

	InspectorIds	[]objectid.ObjectID	`bson:"inspector"`
	StartTime		time.Time			`bson:"startTime"`
	FinishTime		time.Time			`bson:"finishTime"`
	DeviceSummary	[]inspectDeviceSummaryItem	`bson:"deviceSummary"`
	UnionSummary	[]unionDeviceSummaryItem	`bson:"unionSummary"`

	Operator		objectid.ObjectID	`bson:"operator"`
}


/*
type Item struct {
	OID  objectid.ObjectID `bson:"_id,omitempty"` // omitempty not working
	Item string
}
*/


func main() {
	TestDocumentationExamples()
}

func TestDocumentationExamples() {

	client, err := mongo.Connect(
		context.Background(),
		"mongodb://localhost:27017",
		nil)

	if err != nil {
		fmt.Println("Error in mongo.Connect")
	}

	/*
	db0 := client.Database("gotest")
	inventory := db0.Collection("inventory")

	cursor0, _ := inventory.Find(context.Background(), bson.NewDocument(bson.EC.String("item", "canvas")))

	itemRead := Item{}
	for cursor0.Next(context.Background()) {
		cursor0.Decode(&itemRead)
		fmt.Println(itemRead)
	}
	*/


	db := client.Database("city_mocha")
	inspectTasks := db.Collection("inspecttasks")

	cursor, _ := inspectTasks.Find(context.Background(), bson.NewDocument(bson.EC.String("identifier", "XJ201810311130820372")))

	inspTask := inspectTask{}

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&inspTask)
		if err != nil {
			fmt.Println("errr")
		}


		fmt.Printf("item: %v", inspTask)

	}
}
