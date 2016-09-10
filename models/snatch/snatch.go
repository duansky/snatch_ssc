package snatch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"snatch_ssc/ioc"
	"snatch_ssc/job"
	"snatch_ssc/sys"

	"github.com/astaxie/beego"
)

type SscData struct {
	No      sys.NullString `json:"no"`      // 期号
	Results sys.NullString `json:"results"` // 结果
	//	DrawTime sys.NullTime   `json:"drawTime"` // 时间
}

type Snatch interface {
	// 采集
	Snatch() (string, error)
	// 解析
	Resolve(content string) []*SscData
}

// 开始采集
func Proccess() error {
	i := 1
	beego.AppConfig.String("job.spec.ssc.snatch")
	job.CreateJob(beego.AppConfig.String("job.spec.ssc.snatch"), func() {
		beego.Info("--------snatch.ssc.cq")
		if obj, ok := ioc.Create("snatch.ssc.cq"); ok {
			sc := obj.(Snatch)
			content, err := sc.Snatch()
			if err != nil {
				log.Fatal(err)
				return
			}

			datas := sc.Resolve(content)
			j, _ := json.Marshal(datas)
			fmt.Println(string(j))
			ioutil.WriteFile(fmt.Sprintf("c:/%d.txt", i), j, 0644)
			i++
		}
	})

	return nil
}
