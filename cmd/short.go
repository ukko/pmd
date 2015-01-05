package cmd

import (
	"time"

	"github.com/codegangsta/cli"

	"github.com/ukko/pmd/alerts"
	"github.com/ukko/pmd/ui"
)

func Short(c *cli.Context) {
	ui.Init()
	ui.NewBar("Short", time.Minute*5, true, func(state byte) {
		if state == ui.STATE_FINISHED {
			alerts.Notify("Go work")
			alerts.Play()
		}
	})
	ui.MainLoop()
}
