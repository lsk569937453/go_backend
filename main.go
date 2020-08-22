package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"go_backend/controller"
	"go_backend/log"
	_ "go_backend/task"
)

func main() {

	i := 0
	c := cron.New(cron.WithSeconds())
	spec := "*/5 * * * * ?"
	_, err := c.AddFunc(spec, func() {
		i++
		log.Info("cron running:%v", i)
	})
	if err != nil {

		log.Error("Site is down: %v\n", err)
	}
	c.Start()
	initController()

}

func initController() {
	gin.DefaultWriter = log.BaseGinLog()
	r := gin.Default()
	//r.Use(gin.LoggerWithWriter(log.BaseGinLog()))
	r.POST("/api/check/task", controller.InitRouter)
	r.GET("/api/search", controller.Search)
	r.POST("/api/task/add", controller.TaskAdd)
	r.GET("/api/task/getAll", controller.TaskGet)
	r.POST("/api/task/getByUserId", controller.TaskGetByUserId)
	r.POST("/api/task/getById", controller.TaskGetById)
	r.POST("/api/taskHistory/getById", controller.TaskHistoryGetById)
	r.Run(":9394") // listen and serve
}
