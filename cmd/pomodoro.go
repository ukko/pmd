package cmd

import (
	"time"

	"github.com/codegangsta/cli"

	"github.com/ukko/pmd/alerts"
	"github.com/ukko/pmd/ui"
)

func Pomodoro(c *cli.Context) {
	ui.Init()
	ui.NewBar("Work", time.Minute*25, true, func(state byte) {
		if state == ui.STATE_FINISHED {
			alerts.Notify("Сделай перерыв")
			alerts.Play()
		}
	})
	ui.MainLoop()
}
