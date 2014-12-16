package cmd

import (
	"time"

	"github.com/codegangsta/cli"
	"github.com/sethgrid/multibar"

	"github.com/ukko/pmd/alerts"
)

func Pomodoro(c *cli.Context) {

	pb, _ := multibar.New()

	cr := pb.MakeBar(100, "pomodoro")

	go pb.Listen()

	func() {
		all := 25 * 60
		for i := 0; i <= all; i++ {
			cr(100 * i / all)
			time.Sleep(1 * time.Second)
		}
	}()

	alerts.Notify("Сделай перерыв")
	alerts.Play()
}
