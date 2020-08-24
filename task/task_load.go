package task

import (
	"github.com/robfig/cron/v3"
	"github.com/valyala/fasthttp"
	"go_backend/dao"
	"go_backend/log"
	"go_backend/redis"
	"go_backend/vojo"
	"strconv"
	"time"
)

var cronJob *cron.Cron

/**
 * 
 * @Description  get all the task and exec the task cron
 * @Date 2:36 下午 2020/8/24
 **/
func init() {

	alltask := dao.GetAllTask() //get All the task
	c := cron.New(cron.WithSeconds())
	c.Start()
	cronJob = c
	for _, s := range alltask {

		//c := cron.New(cron.WithSeconds())
		cron := s.Task_cron
		url := s.Url
		taskId := s.Id
		AddTask(cron, url, taskId)

	}

}
func AddTask(cron string, url string, taskId int) {
	id, err := cronJob.AddFunc(cron, func() {
		dotask(url, taskId)
	})

	if err != nil {
		log.Error("%v", err.Error())
	} else {
		saveToRedis(taskId, id)

	}
}
func DeleteTask(taskId int) {
	stringID:=strconv.Itoa(taskId)
	localTaskId:=redis.Get(stringID)
	if localTaskId==""{
		log.Info("can not find taskID in redis")
		return
	}
	localTaskIdInt,err:=strconv.Atoi(localTaskId)
	if err!=nil{
		log.Error("DeleteTask error",err.Error())
	}else {
		cronJob.Remove(cron.EntryID(localTaskIdInt))
	}
}

/**
save the mysqlID and taskID to redis
*/
func saveToRedis(taskMysqlId int, taskLocalId cron.EntryID) {
	string1 := strconv.Itoa(taskMysqlId)
	string2 := strconv.Itoa(int(taskLocalId))
	redis.Set(string1, string2)
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
		status = -1
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
