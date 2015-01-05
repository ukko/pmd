package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	STATE_RUNNING = iota
	STATE_BREAKED
	STATE_FINISHED
	STATE_UNKNOWN
)

type Progress struct {
	Duration   time.Duration
	State      byte
	StartTime  time.Time
	FinishTime time.Time
	Title      string
	X, Y, W, H int
	OnFinish   func(state byte)
	progress   int64
	ticker     *time.Ticker
}

// Start timer
func (p *Progress) Start() error {
	p.StartTime = time.Now()
	p.FinishTime = p.StartTime.Add(p.Duration)
	p.State = STATE_RUNNING
	p.progress = 0

	p.ticker = time.NewTicker(time.Millisecond * 200)

	go func() {
		for _ = range p.ticker.C {
			now := int64(time.Now().Sub(p.StartTime))
			p.progress = now * 100 / int64(p.Duration)
			if p.progress >= 100 {
				p.State = STATE_FINISHED
				p.Stop()
			}
			p.Redraw()
		}
	}()

	return nil
}

// Stop timer
func (p *Progress) Stop() error {
	p.ticker.Stop()

	if p.OnFinish != nil {
		p.OnFinish(p.State)
	}
	return nil
}

// Redraw timer
func (p *Progress) Redraw() {
	prefix := fmt.Sprintf(" %5s [ ", p.Title)
	pass := "--:--"
	remain := "--:--"

	if p.State == STATE_RUNNING {
		now := time.Now()
		duration := now.Sub(p.StartTime)
		pass = fmt.Sprintf("%02d:%02d", int(duration.Minutes()), int(duration.Seconds())-int(duration.Minutes())*60)

		duration = p.FinishTime.Sub(now)
		remain = fmt.Sprintf("%02d:%02d", int(duration.Minutes()), int(duration.Seconds())-int(duration.Minutes())*60)
	}

	postfix := fmt.Sprintf(" ] %s/%s", pass, remain)
	lenbar := p.W - len(prefix) - len(postfix) - 2
	bar := make([]string, lenbar)

	for i, _ := range bar {
		if float32(p.progress)/float32(100) > float32(i)/float32(lenbar) {
			bar[i] = "="
		} else {
			bar[i] = "-"
		}
	}

	str := fmt.Sprintf("%s%s%s", prefix, strings.Join(bar, ""), postfix)

	for i, s := range str {
		termbox.SetCell(i+p.X, p.Y, s, termbox.ColorWhite, termbox.ColorBlack)
	}

	termbox.Flush()
}
