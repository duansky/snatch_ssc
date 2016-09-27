package gxklsf

import (
	"snatch_ssc/ioc"
	"snatch_ssc/models/snatch/base"
	"snatch_ssc/models/snatch/inter"
	"strings"

	"snatch_ssc/sys"

	"github.com/astaxie/beego"
	"github.com/duansky/goquery"
)

// 彩票控
type GxcaipiaoSnatch struct {
	base.DataProcesserAbs
}

func init() {
	ioc.RegisterObj("snatch.ssc.gxklsf.gxcaipiao", &GxcaipiaoSnatch{base.DataProcesserAbs{Type: "gxklsf", Site: "gxcaipiao"}})
}

// 抓取网页
func (c *GxcaipiaoSnatch) Snatch() (string, error) {
	doc, err := goquery.NewDocument("http://www.gxcaipiao.com.cn/xml/award_09.xml")

	if err != nil {
		beego.Error(err)
		return "", err
	}

	return doc.Html()
}

// 解析网页数据
func (this *GxcaipiaoSnatch) Resolve(content string) (datas []*inter.SscData) {
	datas = make([]*inter.SscData, 0, 10)
	if !sys.HasValue(content) {
		return datas
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))

	if err != nil {
		beego.Error(err)
	}

	doc.Find("p").Each(func(i int, s *goquery.Selection) {

		if no, ok := s.Attr("id"); ok {
			if results, ok := s.Attr("c"); ok {
				data := new(inter.SscData)
				data.No.SetValue(no)
				data.Results.SetValue(sys.ResultsFillZero(results, ","))
				datas = append(datas, data)
			}
		}
	})

	return datas
}
