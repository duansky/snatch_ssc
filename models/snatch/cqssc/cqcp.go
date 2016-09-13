package cqssc

import (
	"encoding/json"
	"reflect"
	"snatch_ssc/ioc"
	"snatch_ssc/models/snatch/inter"
	"strings"

	"snatch_ssc/sys"

	"github.com/astaxie/beego"
	"github.com/duansky/goquery"
)

// 重庆时时彩官网
type CqcpSnatch struct {
}

func init() {
	ioc.Register("snatch.ssc.cq.cqcp", reflect.TypeOf(new(CqcpSnatch)))
}

// 抓取网页
func (this *CqcpSnatch) Snatch() (string, error) {
	doc, err := goquery.NewDocument("http://www.cqcp.net/game/ssc")
	if err != nil {
		beego.Error(err)
		return "", err
	}

	return doc.Html()
}

// 解析网页数据
func (this *CqcpSnatch) Resolve(content string) (datas []*inter.SscData) {
	datas = make([]*inter.SscData, 0, 10)
	if !sys.HasValue(content) {
		return datas
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		beego.Error(err)
	}

	doc.Find("#openlist").Children().Each(func(i int, s *goquery.Selection) {
		data := new(inter.SscData)
		// For each item found
		if i > 0 {
			no, _ := s.Children().Eq(0).Html()
			data.No.SetValue(no)

			results, _ := s.Children().Eq(1).Html()
			data.Results.SetValue(strings.Replace(results, "-", ",", -1))

			datas = append(datas, data)
		}
	})

	return datas
}

func (this *CqcpSnatch) Processing(datas []*inter.SscData) {
	j, _ := json.Marshal(datas)
	beego.Info(string(j))
}
