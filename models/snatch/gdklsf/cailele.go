package gdklsf

import (
	"snatch_ssc/ioc"
	"snatch_ssc/models/snatch/base"
	"snatch_ssc/models/snatch/inter"

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

	beego.Info("=-===content:", content)

	//	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	//	if err != nil {
	//		beego.Error(err)
	//	}

	//	doc.Find("#openlist").Children().Each(func(i int, s *goquery.Selection) {
	//		data := new(inter.SscData)
	//		// For each item found
	//		if i > 0 {
	//			no, _ := s.Children().Eq(0).Html()
	//			data.No.SetValue(no)

	//			results, _ := s.Children().Eq(1).Html()
	//			data.Results.SetValue(strings.Replace(results, "-", ",", -1))

	//			datas = append(datas, data)
	//		}
	//	})

	return datas
}
