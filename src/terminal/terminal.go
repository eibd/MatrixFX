package terminal

import (
	"fmt"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func HideCursor() {
	fmt.Print("\033[?25l")
}

func ShowCursor() {
	fmt.Print("\033[?25h")
}

func GetTerminalSize() (int, int) {
	ws := &winsize{}
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(ws)))
	return int(ws.Row), int(ws.Col)
}

func PrintColumn(col int, column []rune, height int) {
	fmt.Print("\033[32m")
	for row := 0; row < height; row++ {
		fmt.Printf("\033[%d;%dH%s", row+1, col+1, string(column[row]))
	}
	fmt.Print("\033[0m")
}

func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}
