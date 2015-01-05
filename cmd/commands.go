package cmd

import (
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
		Action:    Long,
	},
	{
		Name:      "short",
		ShortName: "s",
		Usage:     "run short break",
		Action:    Short,
	}}
