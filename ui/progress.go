package ui

import (
	"fmt"
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
	progress   int64
	ticker     *time.Ticker
}

// Start timer
func (p *Progress) Start() error {
	p.StartTime = time.Now()
	p.FinishTime = p.StartTime.Add(p.Duration)
	p.State = STATE_RUNNING
	p.progress = 0

	p.ticker = time.NewTicker(time.Second * 1)

	go func() {
		for range p.ticker.C {
			now := int64(time.Now().Sub(p.StartTime))
			p.progress = now * 100 / int64(p.Duration)
			if p.progress >= 100 {
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
	return nil
}

// Redraw timer
func (p *Progress) Redraw() {
	//pattern := " %5s %2d:%2d/%2d:%2d ================================>------------------------------------------ ] 10:03/14:57 "
	prefix := fmt.Sprintf("%5s ", p.Title)
	pass := "--:--"
	remain := "--:--"

	if p.State == STATE_RUNNING {
		now := time.Now()
		duration := now.Sub(p.StartTime)
		pass = fmt.Sprintf("%02d:%02d", int(duration.Minutes()), int(duration.Seconds())-int(duration.Minutes())*60)

		duration = p.FinishTime.Sub(now)
		remain = fmt.Sprintf("%02d:%02d", int(duration.Minutes()), int(duration.Seconds())-int(duration.Minutes())*60)
	}

	postfix := fmt.Sprintf(" %s/%s", pass, remain)
	bar := make([]string, p.W-len(prefix)-len(postfix)-2)

	//for i, _ := range bar {
	//bar[i] = "-"
	////if float32(p.progress)/float32(100) > float32(i)/float32(len(bar)) {
	////bar[i] = "="
	////} else {
	////bar[i] = "-"
	////}
	//}

	str := fmt.Sprintf("%s%s%s", prefix, bar, postfix)

	for i, s := range str {
		termbox.SetCell(i+p.X, p.Y, s, termbox.ColorWhite, termbox.ColorBlack)
	}

	termbox.Flush()
}
