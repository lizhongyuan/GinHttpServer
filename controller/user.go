package controller

import "github.com/gin-gonic/gin"

func getFilterData (ctx *gin.Context) {
	target := ctx.Query("target")
	scale := ctx.Query("scale")

}
