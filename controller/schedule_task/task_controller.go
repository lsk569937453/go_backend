package schedule_task

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
	c.JSON(http.StatusOK, res)
}
func TaskGetByUserId(c *gin.Context) {
	var req vojo.GetTaskByUserIdReq

	err := c.BindJSON(&req)
	var res vojo.BaseRes
	res.Rescode = vojo.NORMAL_RESPONSE_STATUS
	if err == nil {
		tt, err := dao.GetTaskByUserId(&req)
		if err != nil {
			res.Rescode = vojo.ERROR_RESPONSE_STATUS
			res.Message = err.Error()
		} else {

			data, _ := json.Marshal(tt)
			log.Info("%s", string(data))
			res.Message = tt
		}
	} else {
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		res.Message = err.Error()
		log.Error("bind error:%v", err.Error())
	}
	c.JSON(http.StatusOK, res)

}
func TaskGetById(c *gin.Context) {
	var req vojo.GetTaskByIdReq

	err := c.BindJSON(&req)
	var res vojo.BaseRes
	res.Rescode = vojo.NORMAL_RESPONSE_STATUS
	if err == nil {
		tt, err := dao.GetTaskById(&req)
		if err != nil {
			res.Rescode = vojo.ERROR_RESPONSE_STATUS
			res.Message = err.Error()
		} else {

			data, _ := json.Marshal(tt)
			log.Info("%s", string(data))
			res.Message = tt
		}
	} else {
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		res.Message = err.Error()
		log.Error("bind error:%v", err.Error())
	}
	c.JSON(http.StatusOK, res)

}
func TaskAdd(c *gin.Context) {
	var form vojo.TaskInsertReq

	err := c.BindJSON(&form)
	var res vojo.BaseRes
	if err == nil {

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
			res.Rescode = vojo.ERROR_STATUS_PARAM_WRONG
			c.JSON(http.StatusOK, res)
			return

		}
		tt := dao.AddTask(&form)

		err = task.AddTask(form.CronExpression, form.Url, int(tt))
		if err != nil {
			res.Message = err.Error()
			res.Rescode = vojo.ERROR_RESPONSE_STATUS
		} else {
			res.Message = "添加任务成功"

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
 * @Description  update the task
 * @Date 2:53 下午 2020/8/24
 **/
func TaskUpdate(c *gin.Context) {
	var form vojo.TaskUpdateReq
	form.Id = -1

	err := c.BindJSON(&form)
	var res vojo.BaseRes
	res.Rescode = vojo.NORMAL_RESPONSE_STATUS
	if err == nil {
		tt := dao.UpdateTask(form)
		if tt == nil {
			res.Message = "update success"

			task.DeleteTask(form.Id)
			err := task.AddTask(form.CronExpression, form.Url, form.Id)
			if err != nil {
				res.Message = tt.Error()
				res.Rescode = vojo.ERROR_RESPONSE_STATUS
			}
		} else {
			res.Message = tt.Error()
			res.Rescode = vojo.ERROR_RESPONSE_STATUS
		}
	} else {
		res.Message = err.Error()
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())
	}
	c.JSON(http.StatusOK, res)

}
func TaskDelete(c *gin.Context) {
	var req vojo.TaskDelByIdReq

	err := c.BindJSON(&req)
	var res vojo.BaseRes

	if err == nil {

		tt := dao.DelTask(req)

		res.Rescode = vojo.NORMAL_RESPONSE_STATUS

		if tt == nil {
			task.DeleteTask(req.Id)
			res.Message = "delete success"
		} else {
			res.Rescode = vojo.ERROR_RESPONSE_STATUS
			res.Message = tt.Error()
		}
	} else {
		res.Message = err.Error()
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		log.Error("bind error:%v", err.Error())
	}
	c.JSON(http.StatusOK, res)

}
