package tjklsf

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
type TjflcpwSnatch struct {
	base.DataProcesserAbs
}

func init() {
	ioc.RegisterObj("snatch.ssc.tjklsf.tjflcpw", &TjflcpwSnatch{base.DataProcesserAbs{Type: "tjklsf", Site: "tjflcpw"}})
}

// 抓取网页
func (c *TjflcpwSnatch) Snatch() (string, error) {
	doc, err := goquery.NewDocument("http://www.tjflcpw.com/report/k10_jiben_report.aspx")

	if err != nil {
		beego.Error(err)
		return "", err
	}

	return doc.Html()
}

// 解析网页数据
func (this *TjflcpwSnatch) Resolve(content string) (datas []*inter.SscData) {
	datas = make([]*inter.SscData, 0, 10)
	if !sys.HasValue(content) {
		return datas
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))

	if err != nil {
		beego.Error(err)
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {

		if strings.Contains(s.Text(), "CDATA") {
			datastr := sys.MidStr(s.Text(), "window.onload=function(){ ", "tmpDom=document.createDocumentFragment()")

			for _, v := range strings.Split(datastr, ";") {
				data := new(inter.SscData)
				ss := strings.Split(sys.MidStr(v, "table_add_one_tr(", ")"), ",")
				if len(ss) == 2 {
					data.No.SetValue(strings.Replace(ss[0], "\"", "", -1))
					data.Results.SetValue(strings.TrimSpace(strings.Replace(strings.Replace(ss[1], "|", ",", -1), "\"", "", -1)))
					datas = append(datas, data)
				}
			}
		}
	})

	return datas
}
