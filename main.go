package main

import (
	"github.com/gin-gonic/gin"
	"go_backend/controller"
	"go_backend/log"
	"go_backend/midware"
	_ "go_backend/task"
)

func main() {

	initController()

}

func initController() {
	gin.DefaultWriter = log.BaseGinLog()
	r := gin.Default()
	r.LoadHTMLGlob("./resource/dist/*.html")              // 添加入口index.html
	r.LoadHTMLFiles("./resource/dist/static/*/*")         // 添加资源路径
	r.Static("/static", "./resource/dist/static")         // 添加资源路径
	r.StaticFile("/admin/", "./resource/dist/index.html") //前端接口

	r.Use(midware.IpAuthorize())

	//r.Use(gin.LoggerWithWriter(log.BaseGinLog()))
	r.POST("/api/check/task", controller.InitRouter)
	r.GET("/api/search", controller.Search)
	r.POST("/api/task/add", controller.TaskAdd)
	r.GET("/api/task/getAll", controller.TaskGet)
	r.POST("/api/task/getByUserId", controller.TaskGetByUserId)
	r.POST("/api/task/getById", controller.TaskGetById)
	r.POST("/api/task/updateById", controller.TaskUpdate)
	r.POST("/api/task/delById", controller.TaskDelete)
	r.POST("/api/taskHistory/getByTaskId", controller.TaskHistoryGetByTaskId)
	r.Run(":9393") // listen and serve
}
