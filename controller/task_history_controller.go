package controller

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
	error := c.BindJSON(&req)
	if error == nil {

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

		// fmt.Println(res) // 正常输出msg内容
		c.JSON(http.StatusOK, res)

	} else {
		log.Error("bind error:%v", error.Error())
	}

}

/**
 *
 * @Description history get by page
 * @Date 11:28 上午 2020/9/3
 **/
func TaskHistoryGetByPage(c *gin.Context) {
	var req vojo.GetHistoryByPage
	error := c.BindJSON(&req)
	if error == nil {

		tt, err := dao.HistotyGetByPage(&req)
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

		// fmt.Println(res) // 正常输出msg内容
		c.JSON(http.StatusOK, res)

	} else {
		log.Error("bind error:%v", error.Error())
	}

}
