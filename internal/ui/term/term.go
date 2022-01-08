package term

import (
	"conway/internal/ui/command"
	"conway/internal/universe"
	"fmt"
	"os"
	"os/exec"
)

type Context struct {
	// Make this configurable
}

func (c Context) Render(u universe.Universe) {
	clear()
	fmt.Println(u)
}

func (c Context) Command() command.Command {
	return 0
}

// Terminal related helper functions

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
