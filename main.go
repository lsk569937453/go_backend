package main

import (
	"fmt"
	//"github.com/fullstorydev/grpcurl"
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
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("[GIN] %v |%3d| %13v | %15s | %-7s  %#v %s |\"%s\" \n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
			param.Request.UserAgent(),
		)
		//// your custom format
		//return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		//	param.ClientIP,
		//	param.TimeStamp.Format(time.RFC3339),
		//	param.Method,
		//	param.Path,
		//	param.Request.Proto,
		//	param.StatusCode,
		//	param.Latency,
		//	param.Request.UserAgent(),
		//	param.ErrorMessage,
		//)
	}))
	r.LoadHTMLGlob("./resource/dist/*/*.html")               // 添加入口index.html
	r.LoadHTMLFiles("./resource/dist/static/*/*")            // 添加资源路径
	r.Static("/static", "./resource/dist/")                  // 添加资源路径
	r.StaticFile("/admin/", "./resource/dist/pc/index.html") //前端接口

	r.StaticFile("/m/", "./resource/dist/mobile/index.html") //前端接口

	r.Use(midware.IpAuthorize())

	//r.Use(gin.LoggerWithWriter(log.BaseGinLog()))
	r.POST("/api/task/add", controller.TaskAdd)
	r.GET("/api/task/getAll", controller.TaskGet)
	r.POST("/api/task/getByUserId", controller.TaskGetByUserId)
	r.POST("/api/task/getById", controller.TaskGetById)
	r.POST("/api/task/updateById", controller.TaskUpdate)
	r.POST("/api/task/delById", controller.TaskDelete)
	r.POST("/api/taskHistory/getByTaskId", controller.TaskHistoryGetByTaskId)
	r.POST("/api/taskHistory/getByPage", controller.TaskHistoryGetByPage)
	r.POST("/api/grpc/getServiceList", controller.GrpcGetServiceList)
	r.POST("/api/grpc/remoteInvoke", controller.GrpcRemoteInvoke)
	r.GET("/api/db/dbPing", controller.DbPing)
	fmt.Println("the server has started in the 9393")
	r.Run(":9393") // listen and serve
	//	r.RunTLS(":9393", "resource/client.pem", "resource/client.key")

}
