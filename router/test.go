package router


import (
	"github.com/gin-gonic/gin"
	"../controller"
)


func RouterHandler(engine *gin.Engine) {

	// engine.GET("/user/:name", userGetHandler)
	engine.GET("/user/:name", controller.UserGetHandler)

	// engine.GET("/user/:name/:action", userNameActionHandler)
	engine.GET("/user/:name/:action", controller.UserNameActionHandler)

}
