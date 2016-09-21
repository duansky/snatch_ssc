package base

import (
	"encoding/json"
	"snatch_ssc/models"
	"snatch_ssc/models/snatch/inter"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 数据后续处理接口
type DataProcesser interface {
	// 把采集到的数据进行后续处理, 如入库、推送等等
	Processing([]*inter.SscData, string, string)
	// 获取类型和来源
	GetType() (string, string)
}

type DataProcesserAbs struct {
	Type string // ssc种类
	Site string // 抓取来源网站
}

// 数据后续处理
func (this *DataProcesserAbs) Processing(datas []*inter.SscData, t, s string) {
	// 数据入库
	save(datas, t, s)
	// 数据推送
	push(datas, t, s)
}

func (this *DataProcesserAbs) GetType() (string, string) {

	return this.Type, this.Site
}

// 数据入库
func save(datas []*inter.SscData, t, s string) {
	j, _ := json.Marshal(datas)
	beego.Info(string(j))

	o := orm.NewOrm()
	for _, v := range datas {

		// 也可以直接使用对象作为表名
		var d models.Data
		qs := o.QueryTable(d)
		if err := qs.Filter("no", v.No.String()).Filter("type", t).One(&d); err != nil {
			o.Insert(&models.Data{No: v.No.String(), Results: v.Results.String(), Type: t, Site: s})
		} else {
			// 如果不存在s站点的数据则添加s站点
			if !strings.Contains(d.Site, s) {
				d.Site += "," + s
				if _, err := o.Update(&d); err != nil {
					beego.Error(err)
				}
			}
		}
	}
}

// 数据推送
func push(datas []*inter.SscData, t, s string) {
}
