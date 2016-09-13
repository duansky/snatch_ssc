package inter

import "snatch_ssc/sys"

type SscData struct {
	No      sys.NullString `json:"no"`      // 期号
	Results sys.NullString `json:"results"` // 结果
	//	DrawTime sys.NullTime   `json:"drawTime"` // 时间
}

type Snatch interface {
	// 抓取网页
	Snatch() (string, error)
	// 解析网页数据
	Resolve(string) []*SscData
}
