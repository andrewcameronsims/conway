package ui

import (
	"conway/internal/ui/term"
	"conway/internal/universe"
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
		context: term.Context{},
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
