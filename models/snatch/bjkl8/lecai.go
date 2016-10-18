package bjkl8

import (
	"encoding/json"
	"fmt"
	"snatch_ssc/ioc"
	"snatch_ssc/models/snatch/base"
	"snatch_ssc/models/snatch/inter"
	"strings"

	"snatch_ssc/sys"

	"github.com/astaxie/beego"
	"github.com/duansky/goquery"
)

// 彩票控
type LecaiSnatch struct {
	base.DataProcesserAbs
}

func init() {
	ioc.RegisterObj("snatch.ssc.bjkl8.lecai", &LecaiSnatch{base.DataProcesserAbs{Type: "bjkl8", Site: "lecai"}})
}

// 抓取网页
func (c *LecaiSnatch) Snatch() (string, error) {
	doc, err := goquery.NewDocument("http://baidu.lecai.com/lottery/draw/view/543")

	if err != nil {
		beego.Error(err)
		return "", err
	}

	return doc.Html()
}

// 解析网页数据
func (this *LecaiSnatch) Resolve(content string) (datas []*inter.SscData) {
	datas = make([]*inter.SscData, 0, 10)
	if !sys.HasValue(content) {
		return datas
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))

	if err != nil {
		beego.Error(err)
	}

	j := []byte(sys.MidStr(doc.Text(), "var phaseData =", ";"))

	var rootMap map[string]*json.RawMessage
	err = json.Unmarshal(j, &rootMap)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range rootMap {

		var noMap map[string]*json.RawMessage
		err := json.Unmarshal([]byte(*v), &noMap)
		if err != nil {
			fmt.Println(err)
		}

		for k, v2 := range noMap {
			data := new(inter.SscData)
			data.No.SetValue(k) // 期号

			var resultMap map[string]*json.RawMessage
			err := json.Unmarshal([]byte(*v2), &resultMap)
			if err != nil {
				fmt.Println(err)
			}

			var r struct {
				Red []string `json:"red"`
			}
			json.Unmarshal([]byte(*resultMap["result"]), &r)

			data.Results.SetValue(strings.Join(r.Red, ","))
			datas = append(datas, data)
		}

	}

	return datas
}
