package ui

import (
	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"
)

func drawScreen() {
	width, height := termbox.Size()

	for x := 0; x < width; x++ {
		termbox.SetCell(x, 0, '-', termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(x, height-1, '-', termbox.ColorWhite, termbox.ColorBlack)
	}

	for y := 1; y < height-1; y++ {
		termbox.SetCell(0, y, '|', termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(width-1, y, '|', termbox.ColorWhite, termbox.ColorBlack)
	}
}

// Init main screen
func Init(c *cli.Context) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawScreen()
	termbox.Flush()

loop:

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			drawScreen()
			termbox.Flush()
		case termbox.EventMouse:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			drawScreen()
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}

	}
}
