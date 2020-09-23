package controller

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
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	error := c.BindJSON(&req)
	if error == nil {
		//log.Info(form.Name, form.CronExpression)

		tt, err := service.GrpcGetServiceList(req.Url)

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

func GrpcRemoteInvoke(c *gin.Context) {
	var req vojo.GrpcInvokeReq
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	error := c.BindJSON(&req)
	if error == nil {
		//log.Info(form.Name, form.CronExpression)

		tt, err := service.GrpcRemoteInvoke(req.Url, req.ServiceName, req.MethodName, req.ReqJson)
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
