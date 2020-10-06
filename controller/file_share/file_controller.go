package file_share

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go_backend/log"
	"go_backend/service"
	"go_backend/vojo"
	"net/http"
	"strings"
)

func UploadChunk(c *gin.Context) {
	var res vojo.BaseRes
	key, err := service.SaveChunk(c)

	if err != nil {
		res.Message = fmt.Sprintf("UploadFile error:%s", err.Error())
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())

	} else {
		res.Message = key
	}

	c.JSON(http.StatusOK, res)
}
func MergeChunk(c *gin.Context) {
	var res vojo.BaseRes
	err := service.MergeChunk(c)

	if err != nil {
		res.Message = fmt.Sprintf("UploadFile error:%s", err.Error())
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())

	} else {
		res.Message = ""
	}

	c.JSON(http.StatusOK, res)
}
func MergeChunk2(c *gin.Context) {
	var res vojo.BaseRes
	res.Rescode = vojo.NORMAL_RESPONSE_STATUS
	form, err := c.MultipartForm()
	if err != nil {
		res.Message = fmt.Sprintf("MultipartForm error:%s", err.Error())
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())
		c.JSON(http.StatusOK, res)
		return
	}
	clientId := form.Value["clientId"]
	if clientId == nil || len(clientId) == 0 || clientId[0] == "" {
		res.Message = fmt.Sprintf("clientId error:%s", err.Error())
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())
		c.JSON(http.StatusOK, res)
		return
	}

	files := form.File["file"]
	if files == nil || len(files) == 0 {
		errorMessage := "file error,file is null"
		res.Message = errorMessage
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", errorMessage)
		c.JSON(http.StatusOK, res)
		return
	}
	if len(files) > 1 {
		log.Warn("files length is bigger than 1:%v" + err.Error())
	}
	f := files[0]

	key, err := service.SaveFile(c, f, clientId[0])
	if err != nil {
		res.Message = fmt.Sprintf("UploadFile error:%s", err.Error())
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())

	} else {
		res.Message = key
	}

	c.JSON(http.StatusOK, res)
}

func DownloadFile(ctx *gin.Context) {

	err := service.DownloadService(ctx)

	var res vojo.BaseRes
	if err != nil {
		res.Message = fmt.Sprintf("DownloadFile error:%s", err.Error())
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())
		ctx.JSON(http.StatusForbidden, res)
	}

}
func GetClientID(ctx *gin.Context) {
	uu := uuid.NewV4()
	realKey := uu.String()
	realKey = strings.Replace(realKey, "-", "", -1)
	var res vojo.BaseRes
	res.Rescode = 0

	res.Message = realKey
	// fmt.Println(res) // 正常输出msg内容
	ctx.JSON(http.StatusOK, res)
}
func GetFileList(ctx *gin.Context) {
	clientId := ctx.Query("clientId")
	var res vojo.BaseRes
	if clientId == "" {
		errMessage := "DownloadFile error:clientId is null"
		res.Message = "DownloadFile error:clientId is null"
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", errMessage)
		ctx.JSON(http.StatusForbidden, res)
	}
	result, err := service.GetFileList(clientId)
	if err != nil {
		res.Message = fmt.Sprintf("DownloadFile error:%s", err.Error())
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())
	} else {
		res.Message = result
	}
	ctx.JSON(http.StatusOK, res)

}
