package main

import (
	"fmt"
	"go_backend/controller/file_share"
	"go_backend/controller/task_controller_package"

	//"github.com/fullstorydev/grpcurl"
	"github.com/gin-gonic/gin"
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
	}))
	r.LoadHTMLGlob("./resource/dist/*/*.html")               // 添加入口index.html
	r.LoadHTMLFiles("./resource/dist/static/*/*")            // 添加资源路径
	r.Static("/static", "./resource/dist/")                  // 添加资源路径
	r.StaticFile("/admin/", "./resource/dist/pc/index.html") //前端接口

	r.StaticFile("/m/", "./resource/dist/mobile/index.html") //前端接口

	r.Use(midware.IpAuthorize())

	//r.Use(gin.LoggerWithWriter(log.BaseGinLog()))
	r.POST("/api/task/add", task_controller_package.TaskAdd)
	r.GET("/api/task/getAll", task_controller_package.TaskGet)
	r.POST("/api/task/getByUserId", task_controller_package.TaskGetByUserId)
	r.POST("/api/task/getById", task_controller_package.TaskGetById)
	r.POST("/api/task/updateById", task_controller_package.TaskUpdate)
	r.POST("/api/task/delById", task_controller_package.TaskDelete)
	r.POST("/api/taskHistory/getByTaskId", task_controller_package.TaskHistoryGetByTaskId)
	r.POST("/api/taskHistory/getByPage", task_controller_package.TaskHistoryGetByPage)
	r.POST("/api/grpc/getServiceList", task_controller_package.GrpcGetServiceList)
	r.POST("/api/grpc/remoteInvoke", task_controller_package.GrpcRemoteInvoke)
	r.GET("/api/db/dbPing", task_controller_package.DbPing)

	r.POST("/api/shareFile/uploadChunk", file_share.UploadChunk)
	r.POST("/api/shareFile/mergeChunk", file_share.MergeChunk)
	r.POST("/api/shareFile/download-user-file", file_share.DownloadFile)
	r.GET("/api/shareFile/getClientID", file_share.GetClientID)
	r.GET("/api/shareFile/getFileList", file_share.GetFileList)

	fmt.Println("the server has started in the 9393")
	r.Run(":9393") // listen and serve
	//	r.RunTLS(":9393", "resource/client.pem", "resource/client.key")

}
