package job

import (
	"github.com/astaxie/beego"

	"gopkg.in/robfig/cron.v2"
)

// 定时器
var sscJob *cron.Cron

// 为判断定时器是否已创建
var JobKeysMap = make(map[string]cron.EntryID)

// 是否已启动定时器 true:已启动 false:未启动
var isStartedJob bool = false

func GetJob() *cron.Cron {
	if sscJob == nil {
		sscJob = cron.New()
	}
	return sscJob
}

// 启动定时器
func StartJob() {
	job := GetJob()
	if !isStartedJob {
		beego.Info("timetask starting..........")
		job.Start()
		isStartedJob = true
	}
}

func CreateJob(spec string, cmd func()) (cron.EntryID, error) {
	return GetJob().AddFunc(spec, cmd)
}
