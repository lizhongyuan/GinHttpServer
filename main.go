package main


import (
	"github.com/gin-gonic/gin"
	"./router"
)


func main() {

	gin.SetMode(gin.ReleaseMode)
	ginRouter := gin.Default()

	router.RouterHandler(ginRouter)

	ginRouter.Run(":8080")
}

