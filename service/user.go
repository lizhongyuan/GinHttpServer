package service


import (
	"../model"
	// "context"
	// "github.com/mongodb/mongo-go-driver/bson"
	// "log"
	// "github.com/mongodb/mongo-go-driver/options"
	// "fmt"
)


// var userC *mongo.Collection = model.GetUserCollection()


func GetFilterData(target string, scale string) []model.Manager {

	var data []model.Manager
	if target == "manager" && scale == "dealer" {
		data = getAllManager()
	}

	return data
}


func getAllManager() []model.Manager {

	retArr := []model.Manager{}

	/*
	projection := bson.NewDocument(
		bson.EC.Int32("manager", 1),
	)

	// cursor, err := userC.Find(
	cursor, _:= userC.Find(
		context.Background(),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("manager",
				bson.EC.String("roles", "dealers"),
				bson.EC.SubDocumentFromElements("manager", bson.EC.Boolean("$exists", true)),
			),
		),
		options.Find().SetProjection(projection),
	)

	user := model.User{}


	for cursor.Next(context.Background()) {

		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println("errr")
		}

		curItem := model.Manager{}
		curItem.name = cursor.manager.name
		curItem.contacts = cursor.manager.contacts;
		retArr.append(curItem)
	}
	*/

	return retArr
}


/*
exports.getFilterData = async function getFilterData(user, target, scale) {

let data = void 0;

if (target === 'manager' && scale === 'dealer') {
try {
data = await getAllManager();
}
catch (err) {
throw err;
}
}

return data;
};


async function getAllManager() {
const managerArr = await User.find(
{ roles: 'dealers', manager: {$exists: true} },
{ manager: 1 }
);

const retManagerArr = managerArr.map((item) => {

const curItem = {};

curItem.name = item.manager.name ? item.manager.name : '';
curItem.contacts = item.manager.contacts ? item.manager.contacts : '';

return curItem;
});

return retManagerArr;
}
*/