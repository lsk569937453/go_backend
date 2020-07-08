package main

import (
	"go_backend/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/check/task", controller.InitRouter)
	r.GET("/api/search", controller.Search)

	r.Run(":9394") // listen and serve
}
