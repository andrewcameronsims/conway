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
	// TODO: There is a bug here to fix:
	// 1. press play
	// 2. press some other key
	// 3. pause does not work
	// I think it is because the channel is full
	// and cannot receive subsequent pause command
	ch := make(chan command.Command, 1)
	go func() {
		ch <- ui.context.Command()
	}()

	for {
		select {
		case cmd := <-ch:
			if cmd == command.TogglePlay {
				return
			}
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
