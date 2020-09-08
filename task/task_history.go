package task

import (
	"github.com/robfig/cron/v3"
	"go_backend/dao"
	"go_backend/log"
)

func init() {
	c := cron.New(cron.WithSeconds())
	taslCron := "0 0 0/1 * * ? "
	c.AddFunc(taslCron, func() {
		checkAndDelete()
	})

	c.Start()
}
func checkAndDelete() {
	count := dao.HitoryCount()
	//modify the history count and delete the count
	if count > 30000 {
		dao.HitoryDeleteLast(10000)
	}
	log.Info("history count is %d", count)
}
