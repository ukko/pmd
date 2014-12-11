package main

import (
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/cli"
	//notify "github.com/mqu/go-notify"
	"github.com/sethgrid/multibar"
)

func pomodoro(c *cli.Context) {

	pb, _ := multibar.New()

	cr := pb.MakeBar(100, "pomodoro")

	go pb.Listen()

	go func() {
		pb.Println("func")
		all := 15 * 60
		for i := 0; i <= all; i++ {
			cr(i)
			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}
	}()
	fmt.Println("pomodoro", c.Args().First())
}

func main() {
	app := cli.NewApp()

	app.Name = "pmd"
	app.Commands = []cli.Command{
		{
			Name:      "pomodoro",
			ShortName: "p",
			Usage:     "run pomodoro timer",
			Action:    pomodoro,
		},
		{
			Name:      "long",
			ShortName: "l",
			Usage:     "run long break",
			Action: func(c *cli.Context) {
				fmt.Println("long", c.Args().First())
			},
		},
		{
			Name:      "short",
			ShortName: "s",
			Usage:     "run short break",
			Action: func(c *cli.Context) {
				fmt.Println("short", c.Args().First())
			},
		}}

	app.Run(os.Args)

}
