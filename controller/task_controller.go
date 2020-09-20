package controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go_backend/dao"
	"go_backend/log"
	"go_backend/task"
	"go_backend/vojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 *
 * @Description //TODO
 * @Date 2:25 下午 2020/8/24
 * @Param
 * @return
 **/
func TaskGet(c *gin.Context) {
	var res vojo.BaseRes
	res.Rescode = 0

	tt := dao.GetAllTask()
	data, _ := json.Marshal(tt)
	log.Info("%s", string(data))
	res.Message = tt
	// fmt.Println(res) // 正常输出msg内容
	c.JSON(http.StatusOK, res)

}
func TaskGetByUserId(c *gin.Context) {
	var req vojo.GetTaskByUserIdReq
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	error := c.BindJSON(&req)
	if error == nil {
		//log.Info(form.Name, form.CronExpression)

		tt := dao.GetTaskByUserId(&req)

		var res vojo.BaseRes
		res.Rescode = 0
		data, _ := json.Marshal(tt)
		log.Info("%s", string(data))
		res.Message = tt
		// fmt.Println(res) // 正常输出msg内容
		c.JSON(http.StatusOK, res)

	} else {
		log.Error("bind error:%v", error.Error())
	}

}
func TaskGetById(c *gin.Context) {
	var req vojo.GetTaskByIdReq
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	error := c.BindJSON(&req)
	if error == nil {
		//log.Info(form.Name, form.CronExpression)

		tt := dao.GetTaskById(&req)

		var res vojo.BaseRes
		res.Rescode = 0
		data, _ := json.Marshal(tt)
		log.Info("%s", string(data))
		res.Message = tt
		// fmt.Println(res) // 正常输出msg内容
		c.JSON(http.StatusOK, res)

	} else {
		log.Error("bind error:%v", error.Error())
	}

}
func TaskAdd(c *gin.Context) {
	var form vojo.TaskInsertReq
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	error := c.BindJSON(&form)
	if error == nil {
		//log.Info(form.Name, form.CronExpression)

		var res vojo.BaseRes
		//validate := validator.New()
		//err := validate.Struct(form)
		err := form.UserValidator()
		if err != nil {
			errorMessageArray := make([]*vojo.ErrorMessage, 0)
			for _, err := range err.(validator.ValidationErrors) {
				errorStructObj := &vojo.ErrorMessage{}
				errorStructObj.Field = err.Field()
				if err.Field() == "CronExpression" {
					errorStructObj.Message = "cronExpression is null or illegal "
				} else if err.Field() == "Url" {
					errorStructObj.Message = "url is null or not illegal"
				} else {
					errorStructObj.Message = "name is null or not illegal"
				}
				errorMessageArray = append(errorMessageArray, errorStructObj)
			}

			res.Message = errorMessageArray
			res.Rescode = -2
			c.JSON(http.StatusOK, res)
			return

		}
		tt := dao.AddTask(form)

		task.AddTask(form.CronExpression, form.Url, int(tt))

		res.Message = "添加任务成功"

		//var res vojo.CheckTaskRes
		//res.ResponseCode = 0
		//res.Message = "添加任务成功"
		c.JSON(http.StatusOK, res)

	} else {
		log.Error("bind error:%v", error.Error())
	}

}

/**
 *
 * @Description  update the task
 * @Date 2:53 下午 2020/8/24
 **/
func TaskUpdate(c *gin.Context) {
	var form vojo.TaskUpdateReq
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	error := c.BindJSON(&form)
	if error == nil {
		//log.Info(form.Name, form.CronExpression)

		tt := dao.UpdateTask(form)

		var res vojo.BaseRes
		res.Rescode = tt

		if tt == 0 {
			res.Message = "update success"
			//if update the task success,then remove the cron job
			//in the memory,and add  the new cronjob
			task.DeleteTask(form.Id)
			task.AddTask(form.CronExpression, form.Url, form.Id)
		} else {
			res.Message = "update fail"
		}
		c.JSON(http.StatusOK, res)

	} else {
		log.Error("bind error:%v", error.Error())
	}
}
func TaskDelete(c *gin.Context) {
	var req vojo.TaskDelByIdReq
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	error := c.BindJSON(&req)
	if error == nil {
		//log.Info(form.Name, form.CronExpression)

		tt := dao.DelTask(req)

		var res vojo.BaseRes
		res.Rescode = tt

		if tt == 0 {
			task.DeleteTask(req.Id)
			res.Message = "删除成功"
		} else {
			res.Message = "删除失败"
		}
		//var res vojo.CheckTaskRes
		//res.ResponseCode = 0
		//res.Message = "添加任务成功"
		c.JSON(http.StatusOK, res)

	} else {
		log.Error("bind error:%v", error.Error())
	}
}
