package main


import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/test", getting)

	router.Run(":8080")
}


func getting(ctx *gin.Context) {
	name := ctx.Query("name")
	ctx.String(http.StatusOK, "Hello, %s", name)
}
