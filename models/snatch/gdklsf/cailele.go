package gdklsf

import (
	"fmt"
	"snatch_ssc/ioc"
	"snatch_ssc/models/snatch/base"
	"snatch_ssc/models/snatch/inter"
	"strings"

	"snatch_ssc/sys"

	"github.com/astaxie/beego"
	"github.com/duansky/goquery"
)

// 彩乐乐官网
type CaileleSnatch struct {
	base.DataProcesserAbs
}

func init() {
	ioc.RegisterObj("snatch.ssc.gdklsf.cailele", &CaileleSnatch{base.DataProcesserAbs{Type: "gdklsf", Site: "cailele"}})
}

// 抓取网页
func (c *CaileleSnatch) Snatch() (string, error) {
	doc, err := goquery.NewDocument("http://kjh.cailele.com/kj_klsf.shtml")
	if err != nil {
		beego.Error(err)
		return "", err
	}

	return doc.Html()
}

// 解析网页数据
func (this *CaileleSnatch) Resolve(content string) (datas []*inter.SscData) {
	datas = make([]*inter.SscData, 0, 10)
	if !sys.HasValue(content) {
		return datas
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		beego.Error(err)
	}

	date := doc.Find(".cz_name_period").Text()[:8]
	doc.Find(".stripe").Each(func(i int, s *goquery.Selection) {
		s.Find("tbody tr").Each(func(j int, tr *goquery.Selection) {
			data := new(inter.SscData)
			data.No.SetValue(fmt.Sprintf("%s0%s", date, tr.Children().First().Text()))
			data.Results.SetValue(tr.Children().Last().Text())
			datas = append(datas, data)
		})
	})

	return datas
}
