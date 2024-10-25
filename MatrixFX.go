package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	height = 2
)

var (
	width       = os.Getpagesize()
	coefficient = width / 20
)

func main() {
	rand.Seed(time.Now().UnixNano())
	hideCursor()
	defer showCursor()

	for {
		matrix := generateMatrix()
		printMatrix(matrix)
		shiftMatrixDown(matrix)
		time.Sleep(200 * time.Millisecond)
	}
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func generateMatrix() [][]rune {
	matrix := make([][]rune, height)
	for i := 0; i < height; i++ {
		row := make([]rune, width)
		for j := 0; j < coefficient; j++ {
			if rand.Intn(10) > 2 {
				row[j] = ' '
			} else {
				row[j] = rune(rand.Intn(2) + 48)
			}
		}
		matrix[i] = row
	}
	return matrix
}

func printMatrix(matrix [][]rune) {
	fmt.Print("\033[32m")
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(string(matrix[i][j]))
		}
		fmt.Println()
	}
	fmt.Print("\033[0m")
}

func shiftMatrixDown(matrix [][]rune) {
	for col := 0; col < width; col++ {
		for row := height - 1; row > 0; row-- {
			matrix[row][col] = matrix[row-1][col]
		}
		matrix[0][col] = rune(rand.Intn(94) + 33)
	}
}
