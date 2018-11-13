package main


import (
	"github.com/gin-gonic/gin"
	"./router"
)


func main() {

	gin.SetMode(gin.ReleaseMode)
	ginRouter := gin.Default()

	router.UserHandler(ginRouter)
	router.RouterHandler(ginRouter)

	ginRouter.Run(":8080")
}

