package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"go_backend/controller"
	"go_backend/log"
)

func main() {


	i := 0
	c := cron.New(cron.WithSeconds())
	spec := "*/5 * * * * ?"
	_, err := c.AddFunc(spec, func() {
		i++
		//log("cron running:", i)
	})
	if err != nil {

		log.Error("Site is down: %v\n", err)
	}
	c.Start()
	r := gin.Default()
	r.POST("/api/check/task", controller.InitRouter)
	r.GET("/api/search", controller.Search)
	r.GET("/api/task/getAll", controller.TaskGet)
	r.GET("/api/taskHistory/getAll", controller.TaskHistoryGet)
	r.Run(":9394") // listen and serve
}
