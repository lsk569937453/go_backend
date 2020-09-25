package task_controller_package

import (
	"github.com/gin-gonic/gin"
	"go_backend/dao"
	"go_backend/vojo"
	"net/http"
)

func DbPing(c *gin.Context) {
	var res vojo.BaseRes
	res.Rescode = 0
	var message string
	err := dao.CronDb.Ping()
	if err != nil {
		message = err.Error()
		res.Rescode = -1
	}
	res.Message = message
	// fmt.Println(res) // 正常输出msg内容
	c.JSON(http.StatusOK, res)

}
