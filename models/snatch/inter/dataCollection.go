package inter

import "gopkg.in/robfig/cron.v2"

type DataCollection interface {
	DoCollection() (map[string]cron.EntryID, error)
}
