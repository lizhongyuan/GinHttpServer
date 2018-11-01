package main


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)


func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/test", getFunc)

	router.POST("/post", func(ctx *gin.Context){

		id := ctx.Query("id")
		page := ctx.DefaultQuery("page", "0")
		name := ctx.PostForm("name")
		message := ctx.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)

		ctx.String(http.StatusOK, "Hello, %s", name)

		// id := ctx.Query("id")
		// ctx.String(http.StatusOK, "Hello, %s", id)
	})

	router.Run(":8080")
}


func getFunc(ctx *gin.Context) {
	id := ctx.Query("name")
	ctx.String(http.StatusOK, "wow, %s", id)
}

