package model

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

/*
type manager struct {
	name		string			`bson:"name"`
	contacts	string			`bson:"contacts"`
}
*/

type yunpian struct {
	alarm_tpl_id	string		`bson:"alarm_tpl_id"`
	ok_tpl_id		string		`bson:"ok_tpl_id"`
}

type config struct {
	businessLimit	int			`bson:"businessLimit"`
	businessCount	int			`bson:"businessCount"`
	yunpian			yunpian		`bson:"yunpian"`
}

type character struct {
	fullName		string		`bson:"fullName"`
	shortName		string		`bson:"shortName"`
}

/*
type PolicyType int32
const (
	Policy_MIN      PolicyType = 0
	Policy_MAX      PolicyType = 1
	Policy_MID      PolicyType = 2
	Policy_AVG      PolicyType = 3
)
*/

/*
type roles struct {
	type		string			`bons:"type"`

}
*/

type UserTmp struct {
	OID				objectid.ObjectID		`json:"id" bson:"_id"`
	Nickname		string					`json:"nickname" bson:"nickname"`
	Area			string					`json:"area" bson:"area"`
	City			string					`bson:"city"`
	Province		string					`bson:"province"`
	Contacts		string					`bson:"contacts"`
	ContactsName	string					`bson:"contactsName"`
	Manager			string					`bson:"manager"`
	AppId			string					`bson:"appId"`
	AppKey			string					`bson:"appKey"`
	AppSecret		string					`bson:"appSecret"`
	Roles			string					`bson:"roles"`
	AddBy			string					`bson:"addBy"`
	Cases			string					`bson:"cases"`
	GrantRoles		objectid.ObjectID		`bson:"grantRoles"`
	// Grants

	IndexTags		[]string				`bson:"indexTags"`
	Config			config					`bson:"config"`
	Character		character				`bson:"character"`
}
