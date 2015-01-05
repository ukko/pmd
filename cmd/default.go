package cmd

import "github.com/codegangsta/cli"

// Default action
func Default(c *cli.Context) {
	Pomodoro(c)
}
