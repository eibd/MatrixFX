package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	height = 1
)

var (
	width       = os.Getpagesize()
	coefficient = width / 20
	hiragana    = []rune("ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん")
	katakana    = []rune("ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ")
	kanji       = []rune("一二三四五六七八九十百千万山川田人心力日月火水木金土王天気花雨")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	hideCursor()
	defer showCursor()

	for {
		matrix := generateMatrix()
		printMatrix(matrix)
		shiftMatrixDown(matrix)
		time.Sleep(150 * time.Millisecond)
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
				row[j] = randomJapaneseCharacter()
			}
		}
		matrix[i] = row
	}
	return matrix
}

func randomJapaneseCharacter() rune {
	set := rand.Intn(3)
	switch set {
	case 0:
		return hiragana[rand.Intn(len(hiragana))]
	case 1:
		return katakana[rand.Intn(len(katakana))]
	default:
		return kanji[rand.Intn(len(kanji))]
	}
}

func printMatrix(matrix [][]rune) {
	fmt.Print("\033[32m")
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if rand.Intn(10) > 7 {
				fmt.Printf("\033[1;32m%s\033[0;32m", string(matrix[i][j]))
				fmt.Print(string(matrix[i][j]))
			}
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
		matrix[0][col] = randomJapaneseCharacter()
	}
}
