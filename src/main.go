package main

import (
	"github.com/eibd/MatrixFX/src/animation"
	"github.com/eibd/MatrixFX/src/terminal"
	"os"
	"os/signal"
	"syscall"
)

var done = make(chan bool) // Channel to signal the end of execution

func main() {
	terminal.HideCursor()
	defer terminal.ShowCursor()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGWINCH)
	go func() {
		for range sigChan {
			terminal.ClearScreen()
			animation.MainLoop()
		}
	}()

	animation.MainLoop()
	<-done
}
