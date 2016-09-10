package main

import (
	"snatch_ssc/job"
	"snatch_ssc/models/snatch"
	_ "snatch_ssc/routers"

	"github.com/astaxie/beego"
)

func init() {
	snatch.Proccess()
	job.StartJob()
}

func main() {
	beego.AddTemplateExt("htm")
	beego.Run()
}
