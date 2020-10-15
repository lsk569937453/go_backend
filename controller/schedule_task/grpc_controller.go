package schedule_task

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_backend/log"
	"go_backend/service"
	"go_backend/vojo"
	"net/http"
)

func GrpcGetServiceList(c *gin.Context) {
	var req vojo.GrpcGetServiceListReq

	err := c.BindJSON(&req)
	var res vojo.BaseRes
	res.Rescode = vojo.NORMAL_RESPONSE_STATUS

	if err == nil {

		tt, err := service.GrpcGetServiceList(req.Url)

		if err != nil {
			res.Rescode = vojo.ERROR_RESPONSE_STATUS
			res.Message = err.Error()
		} else {
			data, _ := json.Marshal(tt)
			log.Info("%s", string(data))
			res.Message = tt
		}

	} else {
		log.Error("bind error:%v", err.Error())
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		res.Message = err.Error()

	}
	c.JSON(http.StatusOK, res)

}

func GrpcRemoteInvoke(c *gin.Context) {
	var req vojo.GrpcInvokeReq

	err := c.BindJSON(&req)

	var res vojo.BaseRes
	res.Rescode = vojo.NORMAL_RESPONSE_STATUS
	if err == nil {

		tt, err := service.GrpcRemoteInvoke(req.Url, req.ServiceName, req.MethodName, req.ReqJson)

		if err != nil {
			res.Rescode = vojo.ERROR_RESPONSE_STATUS
			res.Message = err.Error()
		} else {
			data, _ := json.Marshal(tt)
			log.Info("%s", string(data))
			res.Message = tt
		}

	} else {
		log.Error("bind error:%v", err.Error())
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		res.Message = err.Error()
	}
	c.JSON(http.StatusOK, res)

}
