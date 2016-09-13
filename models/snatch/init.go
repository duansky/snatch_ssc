package snatch

import (
	"fmt"
	"snatch_ssc/models/snatch/cqssc"

	"github.com/astaxie/beego"
)

// 开始采集
func Proccess() error {

	/********* 读取配置分别启动CQ、BJ、GX等采集 **********/
	fmt.Println(beego.AppConfig.Strings("snatch::data.collection.item"))

	id, err := cqssc.CreateCqJob()
	_ = id
	if err != nil {
		beego.Error(err)
	}

	return nil
}
