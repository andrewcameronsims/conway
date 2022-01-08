package ui

import (
	"conway/internal/ui/command"
	"conway/internal/universe"
)

type Context interface {
	Render(universe.Universe)
	Command() command.Command
}
