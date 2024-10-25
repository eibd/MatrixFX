package animation

import (
	"github.com/eibd/MatrixFX/src/characters"
	"github.com/eibd/MatrixFX/src/terminal"
	"math/rand"
	"time"
)

const columnSpacing = 3 // Space between columns

func MainLoop() {
	height, width := terminal.GetTerminalSize()
	coefficient := width / columnSpacing
	stopChans := make([]chan bool, coefficient)

	for col := 0; col < coefficient; col++ {
		stopChans[col] = make(chan bool)
		go animateColumn(col*columnSpacing, height, stopChans[col], time.Duration(30+rand.Intn(120))*time.Millisecond)
	}

	select {}
}

func animateColumn(col, height int, stopChan chan bool, interval time.Duration) {
	column := make([]rune, height)
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-stopChan:
			return
		case <-ticker.C:
			for i := height - 1; i > 0; i-- {
				column[i] = column[i-1]
			}
			column[0] = characters.RandomJapaneseCharacterWithSpaces()
			terminal.PrintColumn(col, column, height)
		}
	}
}
