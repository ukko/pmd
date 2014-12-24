package ui

import "time"

//"github.com/nsf/termbox-go"

const (
	STATE_RUNNING = iota
	STATE_BREAKED
	STATE_FINISHED
	STATE_UNKNOWN
)

type Progress struct {
	Title      string
	Duration   time.Duration
	State      byte
	StartTime  time.Time
	FinishTime time.Time
	progress   chan int
}

// Start timer
func (p *Progress) Start() error {
	p.StartTime = time.Now()
	p.FinishTime = p.StartTime.Add(p.Duration)
	p.State = STATE_RUNNING

	return nil
}

// Stop timer
func (p *Progress) Stop() error {
	return nil
}

type ProgressWidget struct {
	Progress
	X, Y, W, H int
}
