package snatch

import (
	"fmt"
	"io/ioutil"
	"snatch_ssc/ioc"
	"snatch_ssc/job"
	"snatch_ssc/models/snatch/base"
	"snatch_ssc/models/snatch/inter"
	"strings"

	"gopkg.in/robfig/cron.v2"

	_ "snatch_ssc/models/snatch/bjkl8"
	_ "snatch_ssc/models/snatch/cqssc"
	_ "snatch_ssc/models/snatch/gdklsf"
	_ "snatch_ssc/models/snatch/gxklsf"
	_ "snatch_ssc/models/snatch/hnklsf"
	_ "snatch_ssc/models/snatch/tjklsf"

	"github.com/astaxie/beego"
)

// 开始采集
func Proccess() error {

	/********* 读取配置分别启动CQ、BJ、GX等采集 **********/
	items := beego.AppConfig.Strings("snatch::data.collection.item")
	beego.Info("采集项:", items)

	for _, item := range items {
		doCollection(item)
	}

	return nil
}

// 采集
func doCollection(t string) (ids map[string]cron.EntryID, err error) {
	sites := beego.AppConfig.Strings(fmt.Sprintf("snatch::data.collection.%s.site", t))
	beego.Info(fmt.Sprintf("====%s site:", t), sites)

	ids = make(map[string]cron.EntryID)
	for _, site := range sites {

		if strings.TrimSpace(site) == "" {
			continue
		}

		key := fmt.Sprintf("snatch.ssc.%s.%s", t, site)
		id, err := job.CreateJob(beego.AppConfig.String(fmt.Sprintf("job::spec.snatch.ssc.%s", t)), func() {
			// 防止意外的错误导致程序退出
			defer func() {
				if err := recover(); err != nil {
					ioutil.WriteFile(fmt.Sprintf("%s_%s.log", t, site), []byte(fmt.Sprintf(fmt.Sprintf("type:%s,site:%s, err:%s", t, site, err))), 777)
				}
			}()

			//			beego.Info("--------", key)
			if obj, ok := ioc.CreateObj(key); ok {
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
		ids[key] = id
	}

	return ids, nil
}
