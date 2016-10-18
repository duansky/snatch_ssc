package bjkl8

import (
	"snatch_ssc/ioc"
	"snatch_ssc/models/snatch/base"
	"snatch_ssc/models/snatch/inter"
	"strings"

	"snatch_ssc/sys"

	"github.com/astaxie/beego"
	"github.com/duansky/goquery"
)

// 北京福彩网
type BwlcSnatch struct {
	base.DataProcesserAbs
}

func init() {
	ioc.RegisterObj("snatch.ssc.bjkl8.bwlc", &BwlcSnatch{base.DataProcesserAbs{Type: "bjkl8", Site: "bwlc"}})
}

// 抓取网页
func (c *BwlcSnatch) Snatch() (string, error) {
	doc, err := goquery.NewDocument("http://www.bwlc.net/bulletin/prevkeno.html")

	if err != nil {
		beego.Error(err)
		return "", err
	}

	return doc.Html()
}

// 解析网页数据
func (this *BwlcSnatch) Resolve(content string) (datas []*inter.SscData) {
	datas = make([]*inter.SscData, 0, 10)
	if !sys.HasValue(content) {
		return datas
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		beego.Error(err)
	}

	table := doc.Find("table:contains('开奖号码')")

	table.Find("tr").Each(func(i int, tr *goquery.Selection) {
		if i > 0 {
			data := new(inter.SscData)
			data.No.SetValue(tr.Children().Eq(0).Text())
			data.Results.SetValue(tr.Children().Eq(1).Text())
			datas = append(datas, data)
		}
	})

	return datas
}
