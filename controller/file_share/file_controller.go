package file_share

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_backend/log"
	"go_backend/service"
	"go_backend/vojo"
	"net/http"
)

func UploadFile(c *gin.Context) {
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

	files := form.File["file"]
	if files == nil || len(files) == 0 {
		res.Message = fmt.Sprintf("file error:%s", err.Error())
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())
		c.JSON(http.StatusOK, res)
		return
	}
	if len(files) > 1 {
		log.Warn("files length is bigger than 1:%v" + err.Error())
	}
	f := files[0]

	key, err := service.SaveFile(c, f)
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
