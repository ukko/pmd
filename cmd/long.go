package cmd

import (
	"time"

	"github.com/codegangsta/cli"

	"github.com/ukko/pmd/alerts"
	"github.com/ukko/pmd/ui"
)

func Long(c *cli.Context) {
	ui.Init()
	ui.NewBar("Long", time.Minute*15, true, func(state byte) {
		if state == ui.STATE_FINISHED {
			alerts.Notify("Сделай перерыв")
			alerts.Play()
		}
	})
	ui.MainLoop()
}
