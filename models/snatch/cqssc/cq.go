package cqssc

import (
	"snatch_ssc/ioc"
	"snatch_ssc/job"
	"snatch_ssc/models/snatch/inter"

	"github.com/astaxie/beego"
	"gopkg.in/robfig/cron.v2"
)

// 重庆时时彩官网
type CqSnatch struct {
}

func CreateCqJob() (cron.EntryID, error) {
	return job.CreateJob(beego.AppConfig.String("job::spec.ssc.snatch"), func() {
		beego.Info("--------snatch.ssc.cq")
		if obj, ok := ioc.Create("snatch.ssc.cq.cqcp"); ok {
			sc := obj.(inter.Snatch)
			content, err := sc.Snatch()
			if err != nil {
				beego.Error(err)
				return
			}
			// 解析html
			datas := sc.Resolve(content)
			DataProcesser := obj.(inter.DataProcesser)
			DataProcesser.Processing(datas)
		}
	})
}
