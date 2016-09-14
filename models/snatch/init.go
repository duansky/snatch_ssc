package snatch

import (
	"fmt"
	"snatch_ssc/ioc"
	"snatch_ssc/models/snatch/inter"

	_ "snatch_ssc/models/snatch/cqssc"

	"github.com/astaxie/beego"
)

// 开始采集
func Proccess() error {

	/********* 读取配置分别启动CQ、BJ、GX等采集 **********/
	items := beego.AppConfig.Strings("snatch::data.collection.item")
	beego.Info("采集项:", items)

	for _, item := range items {
		if obj, ok := ioc.Create(fmt.Sprintf("snatch.ssc.%s", item)); ok {
			dc := obj.(inter.DataCollection)
			dc.DoCollection()
		}
	}

	return nil
}
