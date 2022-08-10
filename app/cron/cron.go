package cron

import (
	"github.com/robfig/cron/v3"
)

var Services *cron.Cron

func init() {
	Services = cron.New(cron.WithSeconds())
	Services.AddFunc("59 59 23 * * *", func() {
		// todo
	})
	Services.Start()
}
