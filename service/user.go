package service

import (
	"github.com/gin-gonic/gin/render"
)

func GetFilterData(user render.JSON, target string, scale string) {

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