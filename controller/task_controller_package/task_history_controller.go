package task_controller_package

import (
	"encoding/json"
	"go_backend/dao"
	"go_backend/log"
	"go_backend/vojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TaskHistoryGetByTaskId(c *gin.Context) {
	var req vojo.GetTaskHistoryByTaskIdReq
	err := c.BindJSON(&req)
	var res vojo.BaseRes
	res.Rescode = vojo.NORMAL_RESPONSE_STATUS
	if err == nil {
		tt, err := dao.HistoryGetById(&req)
		var res vojo.BaseRes
		res.Rescode = vojo.NORMAL_RESPONSE_STATUS
		if err != nil {
			res.Rescode = vojo.ERROR_RESPONSE_STATUS
			res.Message = err.Error()
		} else {
			data, _ := json.Marshal(tt)
			log.Info("%s", string(data))
			res.Message = tt
		}
	} else {
		res.Message = err.Error()
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())
	}
	c.JSON(http.StatusOK, res)

}

/**
 *
 * @Description history get by page
 * @Date 11:28 上午 2020/9/3
 **/
func TaskHistoryGetByPage(c *gin.Context) {
	var req vojo.GetHistoryByPage
	err := c.BindJSON(&req)
	var res vojo.BaseRes
	if err == nil {
		tt, err := dao.HistotyGetByPage(&req)
		res.Rescode = vojo.NORMAL_RESPONSE_STATUS
		if err != nil {
			res.Rescode = vojo.ERROR_RESPONSE_STATUS
			res.Message = err.Error()
		} else {
			data, _ := json.Marshal(tt)
			log.Info("%s", string(data))
			res.Message = tt
		}
	} else {
		res.Message = err.Error()
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())
	}
	c.JSON(http.StatusOK, res)

}
