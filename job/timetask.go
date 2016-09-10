package job

import (
	"fmt"
	"time"

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

func CreateJob(spec string, cmd func()) {
	GetJob().AddFunc(spec, cmd)
}

// 生成定时器
func CreateJob2(prfix string, sourceTime time.Time, handler func(t time.Time)) {
	// 转换成服务器时区
	jobTime := BjToLocal(sourceTime)

	// 闭包
	fn := func(t time.Time) func() {
		return func() {
			handler(t)
		}
	}(sourceTime)

	oHour, oMinute, oSecond := jobTime.Clock()
	spec := fmt.Sprintf("%d %d %d * * ?", oSecond, oMinute, oHour)
	key := prfix + ":" + spec
	if _, ok := JobKeysMap[key]; !ok { // 限制同一时间的定时器只生成一次
		//		log.Info("生成定时器........:", key)
		id, _ := GetJob().AddFunc(spec, fn)
		JobKeysMap[key] = id
	}
}

// 把北京时区转换成服务器时区
func BjToLocal(sourceTime time.Time) time.Time {
	// 转成当前机器时区
	year, month, day := sourceTime.Date()
	hour, min, sec := sourceTime.Clock()
	// 将sourceTime设定为北京时区
	sourceTime = time.Date(year, month, day, hour, min, sec, 0, time.FixedZone("CST", 28800))
	// 转换成服务器时区
	return sourceTime.Local()
}
