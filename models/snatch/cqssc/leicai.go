package cqssc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"snatch_ssc/ioc"
	"snatch_ssc/models/snatch/inter"
	"strings"
	"time"

	"snatch_ssc/sys"

	"github.com/astaxie/beego"
	"github.com/duansky/goquery"
)

// 乐彩官网
type LeicaiSnatch struct {
}

func init() {
	ioc.Register("snatch.ssc.cq.leicai", reflect.TypeOf(new(LeicaiSnatch)))
}

// 抓取网页
func (this *LeicaiSnatch) Snatch() (string, error) {
	header := make(http.Header)
	header.Set("Referer", "http://baidu.lecai.com/lottery/draw/view/200")
	doc, err := goquery.NewDocumentWithHeader("http://baidu.lecai.com/lottery/draw/sorts/ajax_get_draw_data.php?lottery_type=200&date="+time.Now().Format("2006-01-02"), header)
	if err != nil {
		beego.Error(err)
		return "", err
	}
	fmt.Println("===lc:", doc.Text())
	return doc.Text(), nil
}

// 解析网页数据
func (this *LeicaiSnatch) Resolve(content string) (datas []*inter.SscData) {
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

func (this *LeicaiSnatch) Processing(datas []*inter.SscData) {
	j, _ := json.Marshal(datas)
	beego.Info(string(j))
}
