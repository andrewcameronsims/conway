package main

import (
	"conway/internal/ui"
	"conway/internal/universe"
	"fmt"
	"os"
)

const usage = "Usage: conway [file]"

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	src := args[1]
	uni, err := universe.FromFile(src)
	if err != nil {
		fmt.Printf("Error reading source file: %v\n", err)
		os.Exit(1)
	}

	ui := ui.NewTermUI(uni)
	ui.Run()
}
