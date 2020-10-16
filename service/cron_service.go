package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"go_backend/util"
	"go_backend/vojo"
	"sync"
	"time"
)

func GetCronExecResult(c *gin.Context) ([]*vojo.CronResult, error) {
	var req vojo.GetCronExecReq
	err := c.Bind(&req)
	if err != nil {
		return nil, err
	}
	if len(req.CronList) == 0 {
		return nil, fmt.Errorf("cron list array size:0")
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(req.CronList))
	resultChan := make(chan *vojo.CronResult, 100000)

	for _, item := range req.CronList {

		go func(val string) {
			defer func() {
				waitGroup.Done()
			}()
			res, err := newSchedule(val)
			var errMessage string
			if err != nil {
				errMessage = err.Error()
				res = make([]string, 0)
			}
			cronResult := &vojo.CronResult{
				Result:         res,
				CronExpression: val,
				Error:          errMessage,
			}
			resultChan <- cronResult

		}(item)
	}
	waitGroup.Wait()
	res := mergeResult(resultChan)
	return res, nil
}

//calculate the schedule
func newSchedule(cronExpression string) ([]string, error) {
	specParser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	schedule, err := specParser.Parse(cronExpression)
	if err != nil {
		return nil, err
	} else {
		resultList := make([]string, 0)
		tem := time.Now()
		for i := 0; i < 10; i++ {
			next := schedule.Next(tem)
			resultList = append(resultList, util.FormatTime(next))
			tem = next
		}
		return resultList, nil
	}
}
func mergeResult(resultList <-chan *vojo.CronResult) []*vojo.CronResult {
	res := make([]*vojo.CronResult, 0)
	for job := range resultList {
		res = append(res, job)
		if len(resultList) == 0 {
			break
		}
	}
	return res

}
