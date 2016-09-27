package gdklsf

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

	cz_name_period := doc.Find(".cz_name_period").Text()
	var date string
	if len(cz_name_period) < 8 {
		date = time.Now().Format("20060102")
	} else {
		date = doc.Find(".cz_name_period").Text()[:8]
	}

	doc.Find(".stripe").Each(func(i int, s *goquery.Selection) {
		s.Find("tbody tr").Each(func(j int, tr *goquery.Selection) {
			data := new(inter.SscData)
			data.No.SetValue(date + tr.Children().First().Text())
			results := tr.Children().Last().Text()
			data.Results.SetValue(results)

			if !(strings.TrimSpace(results) == "" || strings.TrimSpace(tr.Children().First().Text()) == "") {
				datas = append(datas, data)
			}

		})
	})

	return datas
}
