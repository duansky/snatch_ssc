package snatch

import (
	"log"
	"reflect"
	"snatch_ssc/ioc"
	"strings"

	"snatch_ssc/sys"

	"github.com/PuerkitoBio/goquery"
)

// 重庆时时彩
type CqSnatch struct {
}

func init() {
	ioc.Register("snatch.ssc.cq", reflect.TypeOf(new(CqSnatch)))
}

func (this *CqSnatch) Snatch() (string, error) {
	doc, err := goquery.NewDocument("http://www.cqcp.net/game/ssc")
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return doc.Html()
}

func (this *CqSnatch) Resolve(content string) (datas []*SscData) {
	datas = make([]*SscData, 0, 10)
	if !sys.HasValue(content) {
		return datas
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#openlist").Children().Each(func(i int, s *goquery.Selection) {
		data := new(SscData)
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
