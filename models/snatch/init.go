package snatch

import (
	"snatch_ssc/models/snatch/cqssc"

	"github.com/astaxie/beego"
)

// 开始采集
func Proccess() error {
	id, err := cqssc.CreateCqJob()
	_ = id
	if err != nil {
		beego.Error(err)
	}

	return nil
}
