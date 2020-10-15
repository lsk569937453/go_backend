package schedule_task

import (
	"github.com/gin-gonic/gin"
	"go_backend/service"
	"go_backend/vojo"
	"net/http"
)

func GetCronExecResult(c *gin.Context) {
	cronResultList, err := service.GetCronExecResult(c)
	var res vojo.BaseRes
	if err != nil {
		res.Rescode = vojo.ERROR_RESPONSE_STATUS
		res.Message = err.Error()
	} else {
		res.Message = cronResultList
	}
	c.JSON(http.StatusOK, res)

}
