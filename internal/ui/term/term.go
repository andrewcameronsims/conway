package term

import (
	"conway/internal/ui/command"
	"conway/internal/universe"
	"log"
	"os"
	"os/signal"

	"github.com/gdamore/tcell"
)

type Context struct {
	// Make this configurable
	scr   tcell.Screen
	style tcell.Style
}

func NewContext() *Context {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to initialize terminal screen: %v", err)
	}

	if err := s.Init(); err != nil {
		log.Fatalf("Failed to initialize terminal screen: %v", err)
	}

	style := tcell.StyleDefault
	s.SetStyle(style)
	s.Clear()

	context := &Context{
		scr:   s,
		style: style,
	}

	// Shut down gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			context.Exit()
		}
	}()

	return context
}

func (c Context) Exit() {
	c.scr.Fini()
	os.Exit(0)
}

func (c Context) Render(u universe.Universe) {
	c.scr.Clear()

	for i, row := range u {
		for j, cell := range row {
			if cell {
				c.scr.SetContent(j, i, '█', nil, c.style)
			} else {
				c.scr.SetContent(j, i, '.', nil, c.style)
			}
		}
	}

	c.scr.Show()
}

func (c Context) Command() command.Command {
	ev := c.scr.PollEvent()

	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyLeft:
			return command.Back
		case tcell.KeyRight:
			return command.Forward
		default:
			return command.Exit
		}
	default:
		// eventually we may want to respond to user input
		// that is not a keypress
		return command.None
	}
}
