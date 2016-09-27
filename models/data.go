package models

import "time"

// 数据库模型
type Data struct {
	Id         int
	No         string    // 期号
	Results    string    // 结果
	Type       string    // 种类
	Site       string    // 数据来源网站
	SnatchTime time.Time // 数据获取时间
}
