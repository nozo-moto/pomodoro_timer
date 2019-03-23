package timer

import (
	"time"
)

type Timer struct {
	startTime time.Time
	showTime  chan time.Duration
	CountTime time.Duration
}

func New(countTime time.Duration, showTime chan time.Duration) *Timer {
	return &Timer{
		startTime: time.Time{},
		showTime:  showTime,
		CountTime: countTime,
	}
}

func (t *Timer) Run() {
	timeTicker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timeTicker.C:
			if t.startTime != (time.Time{}) {
				t.showTime <- time.Now().Sub(t.startTime)
			}
		}
	}
}

func (t *Timer) Start() {
	t.startTime = time.Now()
}

func (t *Timer) Stop() {
	t.startTime = time.Time{}
	t.showTime <- 0
}
