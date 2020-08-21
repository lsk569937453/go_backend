package task

import (
	"github.com/robfig/cron/v3"
	"go_backend/dao"
	"github.com/valyala/fasthttp"

)

func init()  {

	alltask:=dao.GetAllTask()
	c := cron.New(cron.WithSeconds())
	//spec := "*/5 * * * * ?"
	//_, err := c.AddFunc(spec, func() {
	//	i++
	//	log.Info("cron running:%v", i)
	//})
	for _,s :=range  alltask{
		cron:=s.Task_cron
		url:=s.Url
		c.AddFunc(cron, func() {

		})
	}

}
func doReq(url string)  {
	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {

	}
}