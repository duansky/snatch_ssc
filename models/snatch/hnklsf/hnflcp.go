package hnklsf

import (
	"snatch_ssc/ioc"
	"snatch_ssc/models/snatch/base"
	"snatch_ssc/models/snatch/inter"
	"strings"
	"time"

	"snatch_ssc/sys"

	"github.com/astaxie/beego"
	"github.com/duansky/goquery"
)

// 彩票控
type HnflcpSnatch struct {
	base.DataProcesserAbs
}

func init() {
	ioc.RegisterObj("snatch.ssc.hnklsf.hnflcp", &HnflcpSnatch{base.DataProcesserAbs{Type: "hnklsf", Site: "hnflcp"}})
}

// 抓取网页
func (c *HnflcpSnatch) Snatch() (string, error) {
	doc, err := goquery.NewDocument("http://www.hnflcp.com/moreKJinfo.asp?cptype=4&inqs=" + time.Now().Format("20060102"))

	if err != nil {
		beego.Error(err)
		return "", err
	}

	return doc.Html()
}

// 解析网页数据
func (this *HnflcpSnatch) Resolve(content string) (datas []*inter.SscData) {
	datas = make([]*inter.SscData, 0, 10)
	if !sys.HasValue(content) {
		return datas
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		beego.Error(err)
	}

	doc.Find("table[width*='98%']").Children().First().Children().Each(func(i int, tr *goquery.Selection) {
		if i > 0 {
			data := new(inter.SscData)
			no := tr.Children().Eq(1).Text()
			if len(no) < 11 {
				no = time.Now().Format("2006") + no
			}
			data.No.SetValue(no) // 期号
			resutls := make([]string, 0, 8)
			tr.Find("tr span").Each(func(j int, span *goquery.Selection) {
				resutls = append(resutls, span.Text())
			})
			data.Results.SetValue(strings.Join(resutls, ","))
			datas = append(datas, data)
		}
	})

	return datas
}
