package task

import (
	"github.com/robfig/cron/v3"
	"github.com/valyala/fasthttp"
	"go_backend/dao"
	"go_backend/log"
	"go_backend/vojo"
	"strconv"
	"time"
)

func init() {

	alltask := dao.GetAllTask()

	//spec := "*/5 * * * * ?"
	//_, err := c.AddFunc(spec, func() {
	//	i++
	//	log.Info("cron running:%v", i)
	//})
	c := cron.New(cron.WithSeconds())
	for _, s := range alltask {

		//c := cron.New(cron.WithSeconds())
		cron := s.Task_cron
		url := s.Url
		taskId := s.Id
		_, err := c.AddFunc(cron, func() {
			dotask(url, taskId)
		})
		if err != nil {
			log.Error("%v", err.Error())
		}

	}
	c.Start()

}
func dotask(url string, taskId int) {
	go func() {
		taskHistory := doReq(url, taskId)
		dao.HistoryInsert(taskHistory)
	}()

}
func doReq(url string, taskId int) vojo.TasksHistory {
	preTime := time.Now()
	status, resp, err := fasthttp.Get(nil, url)
	execTime := time.Since(preTime)
	var responseBody string
	if err != nil {
		responseBody = err.Error()
		log.Error("doReq error:%v", err.Error())
	} else {
		responseBody = string(resp)
	}
	var historyDao vojo.TasksHistory
	historyDao.Exec_code = status
	historyDao.Exec_result = responseBody
	historyDao.Exec_time = strconv.FormatInt(execTime.Milliseconds(), 10)
	historyDao.Task_id = taskId
	return historyDao
}
