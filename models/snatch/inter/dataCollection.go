package inter

import "gopkg.in/robfig/cron.v2"

type DataCollection interface {
	DoCollection() (cron.EntryID, error)
}
