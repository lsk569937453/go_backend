package controller

import (
	"fmt"
	"go_backend/dao"
	"go_backend/vojo"
	"net/http"

	"github.com/gin-gonic/gin"

)

func TaskHistoryGet(c *gin.Context) {
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")

	// if form.Name == "user" && form.CronExpression == "password" {
	// 	c.JSON(200, gin.H{"status": "you are logged in"})
	// } else {
	// 	c.JSON(401, gin.H{"status": "unauthorized"})
	// }
	var res vojo.BaseRes
	res.Rescode = 0

	tt := dao.HistoryGetById()
	fmt.Printf("%s", tt)
	res.Message = tt
	// fmt.Println(res) // 正常输出msg内容
	c.JSON(http.StatusOK, res)

}
