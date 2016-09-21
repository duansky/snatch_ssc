package cqssc

import (
	"encoding/json"
	"net/http"
	"snatch_ssc/ioc"
	"snatch_ssc/models/snatch/base"
	"snatch_ssc/models/snatch/inter"
	"strings"
	"time"

	"snatch_ssc/sys"

	"github.com/astaxie/beego"
	"github.com/duansky/goquery"
)

// 乐彩官网
type LecaiSnatch struct {
	base.DataProcesserAbs
}

func init() {
	ioc.RegisterObj("snatch.ssc.cq.lecai", &LecaiSnatch{base.DataProcesserAbs{Type: "cq", Site: "lecai"}})
}

// 抓取网页
func (this *LecaiSnatch) Snatch() (string, error) {
	header := make(http.Header)
	header.Set("Referer", "http://baidu.lecai.com/lottery/draw/view/200")
	doc, err := goquery.NewDocumentWithHeader("http://baidu.lecai.com/lottery/draw/sorts/ajax_get_draw_data.php?lottery_type=200&date="+time.Now().Format("2006-01-02"), header)
	if err != nil {
		beego.Error(err)
		return "", err
	}

	return doc.Text(), nil
}

// 解析网页数据
func (this *LecaiSnatch) Resolve(content string) (datas []*inter.SscData) {
	datas = make([]*inter.SscData, 0, 10)
	if !sys.HasValue(content) {
		return datas
	}

	var l struct {
		Data struct {
			Data []struct {
				Phase  string `json:"phase"`
				Result struct {
					Result []struct {
						Data []string `json:"data"`
					} `json:"result"`
				} `json:"result"`
			} `json:"data"`
		} `json:"data"`
	}
	json.Unmarshal([]byte(content), &l)

	for _, d := range l.Data.Data {
		data := new(inter.SscData)
		data.No.SetValue(d.Phase[2:])
		data.Results.SetValue(strings.Join(d.Result.Result[0].Data, ","))

		datas = append(datas, data)
	}

	return datas
}

//func (this *LecaiSnatch) Processing(datas []*inter.SscData) {
//	j, _ := json.Marshal(datas)
//	beego.Info(string(j))
//	t, s := this.GetType()
//	beego.Info("===LecaiSnatch_Processing:", t, s)
//}
