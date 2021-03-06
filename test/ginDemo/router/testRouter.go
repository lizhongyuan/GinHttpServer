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


/*
func userGetHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.String(http.StatusOK, "Hello %s", name)
}


func userNameActionHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	action := ctx.Param("action")
	message := name + " is " + action

	ctx.String(http.StatusOK, message)
}
*/
