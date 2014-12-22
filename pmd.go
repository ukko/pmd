package main

import (
	"os"

	"github.com/codegangsta/cli"
	//notify "github.com/mqu/go-notify"
	//"github.com/briandowns/spinner"

	"github.com/ukko/pmd/cmd"
	"github.com/ukko/pmd/ui"
)

func main() {
	app := cli.NewApp()
	app.Name = "pmd"
	app.Version = "0.0.1"
	app.Usage = "Pomodoro timer"
	app.Commands = cmd.Commands
	app.Action = ui.Init
	app.Run(os.Args)
}
