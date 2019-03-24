package timer

import (
	"time"
)

const timerTime = 1 * time.Second

type Timer struct {
	startTime    time.Time
	showTimeChan chan time.Duration
	CountTime    time.Duration
}

func New(countTime time.Duration, showTime chan time.Duration) *Timer {
	return &Timer{
		startTime:    time.Time{},
		showTimeChan: showTime,
		CountTime:    countTime,
	}
}

func (t *Timer) Run() {
	timeTicker := time.NewTicker(timerTime)
	for {
		select {
		case <-timeTicker.C:
			if t.startTime != (time.Time{}) {
				showtime := t.CountTime - time.Now().Sub(t.startTime)
				if showtime < 0 {
					t.Stop()
					break
				}
				t.showTimeChan <- showtime
			}
		}
	}
}

func (t *Timer) Start() {
	t.startTime = time.Now()
}

func (t *Timer) Stop() {
	t.startTime = time.Time{}
	t.showTimeChan <- 0
}

func (t *Timer) Initialize(initalTimer time.Duration) {
	t.showTimeChan <- initalTimer
}
