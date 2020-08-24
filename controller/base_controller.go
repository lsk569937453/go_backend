package controller

import (
	"fmt"
	"go_backend/log"
	"go_backend/vojo"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/iris-contrib/go.uuid"
)

type BaseController struct {
}

func InitRouter(c *gin.Context) {

	form := &vojo.CheckTaskReq{}

	if c.BindJSON(&form) == nil {
		fmt.Println(form.Name, form.CronExpression)

		var res vojo.CheckTaskRes
		res.ResponseCode = 0
		res.Message = "add the task success"
		fmt.Println(res) // 正常输出msg内容
		c.JSON(http.StatusOK, res)

	}
}
func Search(c *gin.Context) {

	var res vojo.BaseRes
	res.Rescode = 0

	slice := make([]vojo.TableRow, 0)

	var i int
	for i = 0; i < 100; i++ {
		var searRes vojo.TableRow
		key := rand.Intn(100000)

		searRes.KeyWord = fmt.Sprintf("%d", key)

		// or error handling
		u2, err := uuid.NewV4()
		if err != nil {
			log.Error("Something went wrong: %s", err)
			return
		}

		searRes.UUID = fmt.Sprintf("%s", u2)
		slice = append(slice, searRes)

	}

	//tt := dao.GetTaskByUserId()
	log.Info("%s", 1)
	res.Message = slice
	// fmt.Println(res) // 正常输出msg内容
	c.JSON(http.StatusOK, res)

}
