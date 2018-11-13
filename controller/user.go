package controller


import (
	"github.com/gin-gonic/gin"
	"../service"
	"net/http"
)


func GetFilterData (ctx *gin.Context) {
	target := ctx.Query("target")
	scale := ctx.Query("scale")

	// name := ctx.Param("name")
	// ctx.String(http.StatusOK, "Hello %s", name)

	data := service.GetFilterData(target, scale)

	ctx.String(http.StatusOK, "%v", data)
}


