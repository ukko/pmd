package ui

import (
	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"
)

const (
	BORDER_VERTICAL            = '│'
	BORDER_HORIZONTAL          = '─'
	BORDER_CORNER_LEFT_TOP     = '┌'
	BORDER_CORNER_RIGHT_TOP    = '┐'
	BORDER_CORNER_LEFT_BOTTOM  = '└'
	BORDER_CORNER_RIGHT_BOTTOM = '┘'
)

func drawScreen() {
	width, height := termbox.Size()

	termbox.SetCell(0, 0, BORDER_CORNER_LEFT_TOP, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(width-1, 0, BORDER_CORNER_RIGHT_TOP, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(0, height-1, BORDER_CORNER_LEFT_BOTTOM, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(width-1, height-1, BORDER_CORNER_RIGHT_BOTTOM, termbox.ColorWhite, termbox.ColorBlack)

	for x := 1; x < width-1; x++ {
		termbox.SetCell(x, 0, BORDER_HORIZONTAL, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(x, height-1, BORDER_HORIZONTAL, termbox.ColorWhite, termbox.ColorBlack)
	}

	for y := 1; y < height-1; y++ {
		termbox.SetCell(0, y, BORDER_VERTICAL, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(width-1, y, BORDER_VERTICAL, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func redrawAll() {
	drawScreen()
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
	redrawAll()
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
			redrawAll()
			termbox.Flush()
		case termbox.EventMouse:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			redrawAll()
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}

	}
}
