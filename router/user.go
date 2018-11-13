package router


import (
	"github.com/gin-gonic/gin"
	"../controller"
)


func UserHandler(engine *gin.Engine) {

	engine.GET("/users/filter", controller.GetFilterData)

}

