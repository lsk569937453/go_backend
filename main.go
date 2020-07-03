package main

import (
	"go_backend/controller"

	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	r.POST("/api/check/task", controller.InitRouter)
	r.Run(":9393") // listen and serve
}
