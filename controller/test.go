package controller


import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func UserGetHandler(ctx *gin.Context) {

	name := ctx.Param("name")

	ctx.String(http.StatusOK, "Hello %s", name)
}


func UserNameActionHandler(ctx *gin.Context) {

	name := ctx.Param("name")
	action := ctx.Param("action")
	message := name + " is " + action

	ctx.String(http.StatusOK, message)

}

