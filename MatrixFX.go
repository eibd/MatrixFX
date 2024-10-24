package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 130
	height = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	for {
		matrix := generateMatrix()
		printMatrix(matrix)
		shiftMatrixDown(matrix)
		time.Sleep(250 * time.Millisecond)
	}
}

func generateMatrix() [][]rune {
	matrix := make([][]rune, height)
	for i := 0; i < height; i++ {
		row := make([]rune, width)
		for j := 0; j < width; j++ {
			if rand.Intn(10) > 2 {
				row[j] = rune(' ')
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
