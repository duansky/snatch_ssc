package main

import (
	"snatch_ssc/job"
	"snatch_ssc/models"
	"snatch_ssc/models/snatch"
	_ "snatch_ssc/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql::jdbc.username")+":"+beego.AppConfig.String("mysql::jdbc.password")+"@tcp("+beego.AppConfig.String("mysql::jdbc.host")+")/ssc?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai", 30)
	orm.RegisterModel(new(models.Data))
	orm.RunSyncdb("default", false, true)

	snatch.Proccess()
	job.StartJob()
}

func main() {
	beego.AddTemplateExt("htm")
	beego.Run()
}
