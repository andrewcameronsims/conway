package ui

import (
	"conway/internal/ui/command"
	"conway/internal/ui/term"
	"conway/internal/universe"
	"time"
)

type UI struct {
	history []universe.Universe
	current int
	context Context
}

func NewTermUI(init universe.Universe) *UI {
	return &UI{
		history: []universe.Universe{init},
		current: 0,
		context: term.NewContext(),
	}
}

func (ui *UI) Run() {
	for {
		ui.Render()
		ui.NextCommand()
	}
}

func (ui *UI) Play() {
	togglePlay := make(chan struct{})
	go func() {
		for {
			cmd := ui.context.Command()
			if cmd == command.TogglePlay {
				togglePlay <- struct{}{}
				break
			}
		}
	}()

	for {
		select {
		case <-togglePlay:
			return
		default:
			ui.Forward(1)
			ui.Render()
			ui.wait()
		}
	}
}

func (ui *UI) wait() {
	// Make this configurable
	time.Sleep(time.Millisecond * 250)
}

func (ui *UI) NextCommand() {
	cmd := ui.context.Command()

	switch cmd {
	case command.TogglePlay:
		ui.Play()
	case command.Back:
		ui.Back(1)
	case command.Forward:
		ui.Forward(1)
	case command.Exit:
		ui.context.Exit()
	}
}

func (ui *UI) Render() {
	universe := ui.history[ui.current]
	ui.context.Render(universe)
}

func (ui *UI) Forward(steps int) {
	ui.current += steps

	for ui.current > len(ui.history)-1 {
		next := ui.history[len(ui.history)-1].Tick()
		ui.history = append(ui.history, next)
	}
}

func (ui *UI) Back(steps int) {
	ui.current -= steps
	if ui.current < 0 {
		ui.current = 0
	}
}
