package task

import (
	"encoding/base64"
	"fmt"
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
	go func() {
		startTask()
	}()

}
func startTask() {
	start := time.Now()

	alltask := dao.GetAllTask() //get All the task
	alltask = FilterTask(alltask)
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
	//记录结束时间
	elapsed := time.Since(start)

	//输出执行时间。
	fmt.Println("App elapsed: ", elapsed)

}

/**
 *
 * @Description  add the task into the memory

 * @Date 2:40 下午 2020/8/31
 **/
func AddTask(cron string, url string, taskId int) {
	id, err := cronJob.AddFunc(cron, func() {
		dotask(url, taskId)
	})

	if err != nil {
		errlog := fmt.Sprintf("AddTask error:%s,taskID:%d,cron:%s,url:%s", err.Error(), taskId, cron, url)
		log.Error("", errlog)
	} else {
		saveToRedis(taskId, id)

	}
}

/**
 *
 * @Description  delete the cron task by taskId
 * @Date 2:54 下午 2020/8/24
 **/
func DeleteTask(taskId int) {
	stringID := strconv.Itoa(taskId)
	localTaskId := redis.Get(stringID)
	if localTaskId == "" {
		log.Info("can not find taskID in redis")
		return
	}
	localTaskIdInt, err := strconv.Atoi(localTaskId)
	if err != nil {
		log.Error("DeleteTask error", err.Error())
	} else {
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

/**
 *
 * @Description  exec the task at the scheduled time
 * @Date 2:41 下午 2020/8/31
 **/
func dotask(url string, taskId int) {

	//new a gorotine to exec the data
	go func() {
		taskHistory := doReq(url, taskId)
		dao.HistoryInsert(taskHistory)
	}()

}

/**
 *
 * @Description  request the url with http get method
 * @Date 11:34 上午 2020/8/25
 **/
func doReq(url string, taskId int) vojo.TasksHistory {
	preTime := time.Now()
	status, resp, err := fasthttp.Get(nil, url)
	execTime := time.Since(preTime)
	var responseBody string
	if err != nil {
		responseBody = err.Error()
		status = -1

		errorlog := fmt.Sprintf("error message:%s,taskId:%d", err.Error(), taskId)
		log.Error("doReq error,%s", errorlog)
	} else {
		responseBody = string(resp)
	}
	base64Res := base64.StdEncoding.EncodeToString([]byte(responseBody))

	var historyDao vojo.TasksHistory
	historyDao.Exec_code = status
	historyDao.Exec_result = base64Res
	historyDao.Exec_time = strconv.FormatInt(execTime.Milliseconds(), 10)
	historyDao.Task_id = taskId
	return historyDao
}

//filter the task by the result
func FilterTask(taskList []vojo.TasksDao) []vojo.TasksDao {
	resultList := make([]vojo.TasksDao, 0)
	for _, item := range taskList {
		url := item.Url
		_, _, err := fasthttp.Get(nil, url)
		if err != nil {
			log.Error("taskID:%d,taskUrl:%s could not send req", item.Id, item.Url)
			dao.UpdateTaskStatusByTaskId(item.Id, vojo.RES_ERROR)
			continue
		} else {
			dao.UpdateTaskStatusByTaskId(item.Id, vojo.RES_NORMAL)

		}
		resultList = append(resultList, item)

	}
	return resultList
}
