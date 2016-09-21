package cqssc

import (
	"reflect"
	"snatch_ssc/ioc"
	"snatch_ssc/job"
	"snatch_ssc/models/snatch/base"
	"snatch_ssc/models/snatch/inter"

	"github.com/astaxie/beego"
	"gopkg.in/robfig/cron.v2"
)

// 重庆时时彩官网
type CqCollection struct {
}

func init() {
	ioc.Register("snatch.ssc.cq", reflect.TypeOf(new(CqCollection)))
}

// 采集
func (this *CqCollection) DoCollection() (ids map[string]cron.EntryID, err error) {
	id, err := job.CreateJob(beego.AppConfig.String("job::spec.ssc.snatch"), func() {
		beego.Info("--------snatch.ssc.cqcp")
		if obj, ok := ioc.CreateObj("snatch.ssc.cq.cqcp"); ok {
			sc := obj.(inter.Snatch)
			content, err := sc.Snatch()
			if err != nil {
				beego.Error(err)
				return
			}
			// 解析html
			datas := sc.Resolve(content)
			dataProcesser := obj.(base.DataProcesser)
			t, s := dataProcesser.GetType()
			dataProcesser.Processing(datas, t, s)
		}
	})
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	ids = make(map[string]cron.EntryID)
	ids["snatch.ssc.cq.cqcp"] = id

	id, err = job.CreateJob(beego.AppConfig.String("job::spec.ssc.snatch"), func() {
		beego.Info("--------snatch.ssc.cqlecai")
		if obj, ok := ioc.CreateObj("snatch.ssc.cq.lecai"); ok {
			sc := obj.(inter.Snatch)
			content, err := sc.Snatch()
			if err != nil {
				beego.Error(err)
				return
			}
			// 解析html
			datas := sc.Resolve(content)
			dataProcesser := obj.(base.DataProcesser)
			t, s := dataProcesser.GetType()
			dataProcesser.Processing(datas, t, s)
		}

	})
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	ids["snatch.ssc.cq.leicai"] = id

	return ids, nil
}
