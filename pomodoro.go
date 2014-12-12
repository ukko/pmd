package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/codegangsta/cli"
	//notify "github.com/mqu/go-notify"
	"github.com/sethgrid/multibar"
)

func pomodoro(c *cli.Context) {

	pb, _ := multibar.New()

	cr := pb.MakeBar(100, "pomodoro")

	go pb.Listen()

	func() {
		all := 15 * 60
		for i := 0; i <= all; i++ {
			cr(100 * i / all)
			time.Sleep(1 * time.Second)
		}
	}()

	cmd := exec.Command("/usr/bin/cvlc", "--play-and-exit", "--no-loop", "./alarm.oga")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
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
