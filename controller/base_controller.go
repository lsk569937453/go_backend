package controller

import (
	"fmt"
	"go_backend/vojo"
	"net/http"

	"github.com/gin-gonic/gin"

)

type BaseController struct {
}

func InitRouter(c *gin.Context) {
	form := &vojo.CheckTaskReq{}
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	if c.BindJSON(&form) == nil {
		fmt.Println(form.Name, form.CronExpression)
		// if form.Name == "user" && form.CronExpression == "password" {
		// 	c.JSON(200, gin.H{"status": "you are logged in"})
		// } else {
		// 	c.JSON(401, gin.H{"status": "unauthorized"})
		// }
		var res vojo.CheckTaskRes
		res.ResponseCode = 0
		res.Message = "添加任务成功"
		fmt.Println(res) // 正常输出msg内容
		c.JSON(http.StatusOK, res)

	}
}
