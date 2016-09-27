package gxklsf

import (
	"net/http"
	"snatch_ssc/ioc"
	"snatch_ssc/models/snatch/base"
	"snatch_ssc/models/snatch/inter"
	"strings"

	"snatch_ssc/sys"

	"github.com/astaxie/beego"
	"github.com/duansky/goquery"
)

// 彩票控
type CaipiaokongSnatch struct {
	base.DataProcesserAbs
}

func init() {
	ioc.RegisterObj("snatch.ssc.gxklsf.caipiaokong", &CaipiaokongSnatch{base.DataProcesserAbs{Type: "gxklsf", Site: "caipiaokong"}})
}

// 抓取网页
func (c *CaipiaokongSnatch) Snatch() (string, error) {
	header := make(http.Header)
	header.Set("Cookie", "BAIDU_SSP_lcr=https://www.baidu.com/link?url=xIoltckitMHkiQFQu4DDEsQ-a8EI0rhgOTG9vfT2gfGcodJzYtvBagP-_QCmFj-fa1lj6TlH5XetCjMIrJmyhq&wd=&eqid=b2e813dd000000930000000657e4ca19; __cfduid=d843b7d264e122df43a139b1d941e47411474595145; caipiaokong_4891_saltkey=QkVR4BwR; caipiaokong_4891_lastvisit=1474591545; caipiaokong_4891_caipiaokong_eNr=1; caipiaokong_4891_lastact=1474612009%09index.php%09gxklsf; Hm_lvt_1fa650cb7d8eae53d0e6fbd8aec3eb67=1474602560,1474602633,1474602647,1474612012; Hm_lpvt_1fa650cb7d8eae53d0e6fbd8aec3eb67=1474612012")
	doc, err := goquery.NewDocumentWithHeader("http://www.caipiaokong.com/lottery/gxklsf.html", header)

	if err != nil {
		beego.Error(err)
		return "", err
	}

	return doc.Html()
}

// 解析网页数据
func (this *CaipiaokongSnatch) Resolve(content string) (datas []*inter.SscData) {
	datas = make([]*inter.SscData, 0, 10)
	if !sys.HasValue(content) {
		return datas
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))

	if err != nil {
		beego.Error(err)
	}

	doc.Find(".dt").Filter("table").Find("tr").Each(func(i int, s *goquery.Selection) {
		data := new(inter.SscData)
		if i > 0 {

			results := make([]string, 0, 8)
			s.Find(".brs").Each(func(j int, brs *goquery.Selection) {
				results = append(results, brs.Text())
			})
			if len(results) > 0 {
				data.No.SetValue(s.Find(".xs0").First().Text()[3:12])
				data.Results.SetValue(strings.Join(results, ","))
				datas = append(datas, data)
			}
		}
	})

	return datas
}
