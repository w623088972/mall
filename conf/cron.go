package conf

import (
	"github.com/robfig/cron/v3"
)

//Cron Cron
type Cron struct {
	Self *cron.Cron
}

//CronC 定时任务系统
var CronC *Cron

//Init 日志初始化
func (c *Cron) Init() {
	CronC = &Cron{
		Self: cron.New(cron.WithSeconds()),
	}
	LOG.Self.Info("Cron Init done")
}
