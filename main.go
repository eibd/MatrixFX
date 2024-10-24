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

	// Configuração para ocultar o cursor e limpar a tela (Cursor hiding and screen clearing)
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	// Gerar o efeito cascata (Generate the cascading effect)
	for {
		matrix := generateMatrix()
		printMatrix(matrix)
		shiftMatrixDown(matrix) // New function to move characters down
		time.Sleep(250 * time.Millisecond)
	}
}

func generateMatrix() [][]rune {
	matrix := make([][]rune, height)
	for i := 0; i < height; i++ {
		row := make([]rune, width)
		for j := 0; j < width; j++ {
			if rand.Intn(10) > 2 { // Probability of a space: 20%
				row[j] = rune(' ')
			} else {
				row[j] = rune(rand.Intn(2) + 48) // Printable characters (33 - 126)
			}
		}
		matrix[i] = row
	}
	return matrix
}

func printMatrix(matrix [][]rune) {
	// Define a cor verde (código ANSI 32) antes de imprimir a matriz
	fmt.Print("\033[32m")
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(string(matrix[i][j]))
		}
		fmt.Println()
	}
	// Reseta a cor após imprimir a matriz
	fmt.Print("\033[0m")
}

func shiftMatrixDown(matrix [][]rune) {
	// Move characters down one row, handling the bottom row
	for col := 0; col < width; col++ {
		for row := height - 1; row > 0; row-- {
			matrix[row][col] = matrix[row-1][col]
		}
		// Fill the top row with random characters again
		matrix[0][col] = rune(rand.Intn(94) + 33)
	}
}
