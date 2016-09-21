package inter

import (
	"snatch_ssc/sys"

	"gopkg.in/robfig/cron.v2"
)

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

// 大采集接口, CQ、BJ、GX实现此接口
type DataCollection interface {
	// 进行抓取、解析...
	DoCollection() (map[string]cron.EntryID, error)
}
