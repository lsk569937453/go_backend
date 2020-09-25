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

	f, err := c.FormFile("share_file")

	var res vojo.BaseRes
	res.Rescode = vojo.NORMAL_RESPONSE_STATUS

	if err != nil {
		res.Message = fmt.Sprintf("UploadFile error:%s", err.Error())
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())

	} else {
		key, err := service.SaveFile(c, f)
		if err != nil {
			res.Message = fmt.Sprintf("UploadFile error:%s", err.Error())
			res.Rescode = vojo.ERROR_RESPONSE_STATUS
			log.Error("bind error:%v", err.Error())

		} else {
			res.Message = key
		}
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
		ctx.JSON(http.StatusOK, res)
	}

}
