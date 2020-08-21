package controller

import (
	"encoding/json"
	"go_backend/dao"
	"go_backend/log"
	"go_backend/vojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TaskGet(c *gin.Context) {
	var res vojo.BaseRes
	res.Rescode = 0

	tt := dao.GetAllTask()
	data, _ := json.Marshal(tt)
	log.Info("%s",string(data))
	res.Message = tt
	// fmt.Println(res) // 正常输出msg内容
	c.JSON(http.StatusOK, res)

}
func TaskGetByUserId(c *gin.Context) {
	var req  vojo.GetTaskByUserIdReq
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	error:=c.BindJSON(&req)
	if error == nil {
		//log.Info(form.Name, form.CronExpression)

		tt := dao.GetTaskByUserId(&req)

		var res vojo.BaseRes
		res.Rescode = 0
		data, _ := json.Marshal(tt)
		log.Info("%s",string(data))
		res.Message = tt
		// fmt.Println(res) // 正常输出msg内容
		c.JSON(http.StatusOK, res)


	}else {
		log.Error("bind error:%v",error.Error())
	}

}
func TaskGetById(c *gin.Context) {
	var req  vojo.GetTaskByIdReq
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	error:=c.BindJSON(&req)
	if error == nil {
		//log.Info(form.Name, form.CronExpression)

		tt := dao.GetTaskById(&req)

		var res vojo.BaseRes
		res.Rescode = 0
		data, _ := json.Marshal(tt)
		log.Info("%s",string(data))
		res.Message = tt
		// fmt.Println(res) // 正常输出msg内容
		c.JSON(http.StatusOK, res)


	}else {
		log.Error("bind error:%v",error.Error())
	}

}
func TaskAdd(c *gin.Context) {
	var form  vojo.TaskInsertReq
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	error:=c.BindJSON(&form)
	if error == nil {
		//log.Info(form.Name, form.CronExpression)

		tt := dao.AddTask(form)

		tt.Message="添加任务成功"

		//var res vojo.CheckTaskRes
		//res.ResponseCode = 0
		//res.Message = "添加任务成功"
		c.JSON(http.StatusOK, tt)

	}else {
		log.Error("bind error:%v",error.Error())
	}

}
