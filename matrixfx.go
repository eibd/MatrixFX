package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

// Constants for column spacing
const (
	columnSpacing = 3 // Space between columns
)

// Predefined Japanese character sets as runes
var (
	hiragana = []rune("ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん")
	katakana = []rune("ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ")
	kanji    = []rune("一二三四五六七八九十百千万山川田人心力日月火水木金土王天気花雨")
)

// Struct to hold terminal size information
type winsize struct {
	Row    uint16 // Number of rows
	Col    uint16 // Number of columns
	Xpixel uint16 // Horizontal pixel size
	Ypixel uint16 // Vertical pixel size
}

var done = make(chan bool) // Channel to signal the end of execution

func main() {
	hideCursor()       // Hide the terminal cursor
	defer showCursor() // Ensure the cursor is shown again at the end

	// Set up a channel to capture terminal resize signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGWINCH)

	go func() {
		for range sigChan {
			// Clear the screen and redraw the matrix when the terminal is resized
			clearScreen()
			mainLoop()
		}
	}()

	// Start the main loop
	mainLoop()

	<-done // Keep the program running
}

// Hide the terminal cursor
func hideCursor() {
	fmt.Print("\033[?25l")
}

// Show the terminal cursor
func showCursor() {
	fmt.Print("\033[?25h")
}

// Main loop for generating animations
func mainLoop() {
	height, width := getTerminalSize() // Get the terminal size

	// Calculate the number of columns to generate, respecting the spacing
	coefficient := width / columnSpacing

	// Create a list of channels for controlling goroutines
	stopChans := make([]chan bool, coefficient)

	// Start column animations in goroutines
	for col := 0; col < coefficient; col++ {
		stopChans[col] = make(chan bool) // Create a channel for each column
		// Start a goroutine for each column with a random interval
		go animateColumn(col*columnSpacing, height, stopChans[col], time.Duration(30+rand.Intn(120))*time.Millisecond)
	}

	// Wait for the next signal or event
	select {}
}

// Get the current size of the terminal
func getTerminalSize() (int, int) {
	ws := &winsize{}
	// Use a syscall to get the window size
	syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))
	return int(ws.Row), int(ws.Col)
}

// Animate a column in the terminal
func animateColumn(col, height int, stopChan chan bool, interval time.Duration) {
	column := make([]rune, height) // Create a column buffer

	// Initial random delay for synchronization
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)

	ticker := time.NewTicker(interval) // Create a ticker for this column with a variable interval
	defer ticker.Stop()                // Ensure the ticker is stopped when done

	for {
		select {
		case <-stopChan: // Stop if signaled
			return
		case <-ticker.C: // Execute every tick
			// Shift characters down in the column
			for i := height - 1; i > 0; i-- {
				column[i] = column[i-1]
			}
			// Add a new character at the top of the column
			column[0] = randomJapaneseCharacterWithSpaces()

			// Print the updated column
			printColumn(col, column, height)
		}
	}
}

// Generate a random Japanese character with a chance of a space
func randomJapaneseCharacterWithSpaces() rune {
	if rand.Intn(10) > 2 { // 70% chance to return a space
		return ' '
	}
	return randomJapaneseCharacter() // Return a random Japanese character
}

// Generate a random Japanese character from one of the three sets
func randomJapaneseCharacter() rune {
	set := rand.Intn(3) // Randomly choose a character set
	switch set {
	case 0:
		return hiragana[rand.Intn(len(hiragana))] // Return a random hiragana character
	case 1:
		return katakana[rand.Intn(len(katakana))] // Return a random katakana character
	default:
		return kanji[rand.Intn(len(kanji))] // Return a random kanji character
	}
}

// Print a column in the specified position
func printColumn(col int, column []rune, height int) {
	fmt.Print("\033[32m") // Set text color to green
	for row := 0; row < height; row++ {
		// Move the cursor to the specified row and column, and print the character
		fmt.Printf("\033[%d;%dH%s", row+1, col+1, string(column[row]))
	}
	fmt.Print("\033[0m") // Reset text formatting
}

// Clear the terminal screen
func clearScreen() {
	fmt.Print("\033[2J\033[H") // ANSI escape codes to clear the screen and move cursor to home position
}
