package main

import (
	"fmt"
	"go_backend/controller/file_share"
	"go_backend/controller/schedule_task"

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
	r.POST("/api/task/add", schedule_task.TaskAdd)
	r.GET("/api/task/getAll", schedule_task.TaskGet)
	r.POST("/api/task/getByUserId", schedule_task.TaskGetByUserId)
	r.POST("/api/task/getById", schedule_task.TaskGetById)
	r.POST("/api/task/updateById", schedule_task.TaskUpdate)
	r.POST("/api/task/delById", schedule_task.TaskDelete)
	r.POST("/api/taskHistory/getByTaskId", schedule_task.TaskHistoryGetByTaskId)
	r.POST("/api/taskHistory/getByPage", schedule_task.TaskHistoryGetByPage)
	r.POST("/api/grpc/getServiceList", schedule_task.GrpcGetServiceList)
	r.POST("/api/grpc/remoteInvoke", schedule_task.GrpcRemoteInvoke)
	r.GET("/api/db/dbPing", schedule_task.DbPing)

	r.POST("/api/shareFile/uploadChunk", file_share.UploadChunk)
	r.POST("/api/shareFile/mergeChunk", file_share.MergeChunk)
	r.POST("/api/shareFile/download-user-file", file_share.DownloadFile)
	r.POST("/api/shareFile/downloadChunk", file_share.DownloadChunk)
	r.GET("/api/shareFile/getClientID", file_share.GetClientID)
	r.GET("/api/shareFile/getFileList", file_share.GetFileList)
	r.POST("/api/cron/getCronExecResult", schedule_task.GetCronExecResult)

	fmt.Println("the server has started in the 9393")
	r.Run(":9393") // listen and serve
	//	r.RunTLS(":9393", "resource/client.pem", "resource/client.key")

}
