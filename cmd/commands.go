package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	{
		Name:      "pomodoro",
		ShortName: "p",
		Usage:     "run pomodoro timer",
		Action:    Pomodoro,
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
