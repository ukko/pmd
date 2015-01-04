package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ukko/pmd/ui"
)

// Default action
func Default(c *cli.Context) {
	ui.Init()
}
