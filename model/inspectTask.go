package model


import (
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

type InspectTask struct {
	OID			objectid.ObjectID				`json:"id" bson:"_id,omitempty"`
	Name		string							`json:"name" bson:"name"`
	Identifier	string							`json:"identifier" bson:"identifier"`

	Status			int							`json:"status" bson:"status"`
	CreatedBy		objectid.ObjectID			`bson:"createdBy"`
	CreatedTime		time.Time					`bson:"createdTime"`
	BeginTime		time.Time					`bson:"beginTime"`
	EndTime			time.Time					`bson:"endTime"`

	InspectorIds	[]objectid.ObjectID			`bson:"inspector"`
	StartTime		time.Time					`bson:"startTime"`
	FinishTime		time.Time					`bson:"finishTime"`
	DeviceSummary	[]inspectDeviceSummaryItem	`bson:"deviceSummary"`
	UnionSummary	[]unionDeviceSummaryItem	`bson:"unionSummary"`

	Operator		objectid.ObjectID			`bson:"operator"`
}
