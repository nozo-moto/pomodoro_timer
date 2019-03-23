package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/nozo-moto/pomodoro_timer/timer"
)

var (
	showTimeLabel widget.Label
)

const (
	countTime = time.Minute * 25
)

func init() {
	showTimeLabel.Alignment = fyne.TextAlignCenter
}

func main() {
	showTimeChan := make(chan time.Duration)
	timer := timer.New(countTime, showTimeChan)
	go timer.Run()
	go func() {
		for {
			select {
			case showtime := <-showTimeChan:
				showTimeLabel.SetText(formatTime(showtime))
			}
		}
	}()
	startButton := &widget.Button{
		Text: "Start", OnTapped: func() { timer.Start() },
	}
	stopButton := &widget.Button{
		Text: "Stop", OnTapped: func() { timer.Stop() },
	}

	a := app.New()
	w := a.NewWindow("Timer")
	w.Resize(fyne.Size{Width: 150, Height: 150})

	canvasObjects := []fyne.CanvasObject{
		&widget.Label{Text: "Timer", Alignment: fyne.TextAlignCenter},
		&showTimeLabel,
		startButton,
		stopButton,
	}

	w.SetContent(&widget.Box{Children: canvasObjects})
	w.ShowAndRun()
}

func formatTime(t time.Duration) string {
	h := t / time.Hour
	m := (t - h*time.Hour) / time.Minute
	s := (t - h*time.Hour - m*time.Minute) / time.Second

	return fmt.Sprintf("%d:%d:%d", h, m, s)
}
