package main

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

type boardElement int

const (
	none    boardElement = 0
	red     boardElement = 1
	green   boardElement = 2
	yellow  boardElement = 3
	blue    boardElement = 4
	magenta boardElement = 5
	cyan    boardElement = 6
	white   boardElement = 7
)

func boardElementToString(element boardElement) string {
	switch element {
	case none:
		return " "
	case red:
		return "\033[41m \033[0m"
	case green:
		return "\033[42m \033[0m"
	case yellow:
		return "\033[43m \033[0m"
	case blue:
		return "\033[44m \033[0m"
	case magenta:
		return "\033[45m \033[0m"
	case cyan:
		return "\033[46m \033[0m"
	case white:
		return "\033[47m \033[0m"
	}
	return ""
}

func main() {
	prepareTerminal()

	board := createContainer(10)
	blockA := createContainer(5)
	blockB := createContainer(5)
	blockC := createContainer(5)

	for i := 1; i < 8; i++ {
		board[i][i] = boardElement(i)

		if i == 2 {
			blockA = blockShape(0)
		}
		boardDrawing := drawBoard(board)
		blockADrawing := drawBoard(blockA)
		blockBDrawing := drawBoard(blockB)
		blockCDrawing := drawBoard(blockC)

		display := mergeBoardsHorizontally(" ", boardDrawing, blockADrawing, blockBDrawing, blockCDrawing)

		fmt.Printf("\033[0;0H")
		fmt.Print(display)
		time.Sleep(1 * time.Second)
	}
}

func createContainer(size int) [][]boardElement {
	board := make([][]boardElement, size)
	for i := 0; i < size; i++ {
		board[i] = make([]boardElement, size)
	}
	return board
}

func blockShape(number int) [][]boardElement {
	switch number {
	case 0:
		return [][]boardElement{
			{red, none, none, none, none},
			{none, none, none, none, none},
			{none, none, none, none, none},
			{none, none, none, none, none},
			{none, none, none, none, none}}

	case 1:
		return [][]boardElement{
			{green, none, none, none, none},
			{green, none, none, none, none},
			{green, green, green, none, none},
			{none, none, none, none, none},
			{none, none, none, none, none}}
	default:
		return [][]boardElement{
			{none, none, none, none, none},
			{none, none, none, none, none},
			{none, none, none, none, none},
			{none, none, none, none, none},
			{none, none, none, none, none}}
	}
}

func prepareTerminal() {
	width, height, err := terminal.GetSize(0)
	if err != nil {
		panic("Terminal error")
	}
	fmt.Printf("\033[0;0H")
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func drawBoard(board [][]boardElement) string {
	s := "\u250F"
	for x := 0; x < len(board[0]); x++ {
		s += "\u2501"
	}
	s += "\u2513\n"
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if y == 0 {
				s += "\u2503"
			}
			s += boardElementToString(board[x][y])
			if y == len(board[x])-1 {
				s += "\u2503"
			}
		}
		s += "\n"
	}
	s += "\u2517"
	for x := 0; x < len(board[0]); x++ {
		s += "\u2501"
	}
	s += "\u251B\n"
	return s
}

func mergeBoardsHorizontally(separator string, boards ...string) string {
	merged := ""
	maxLines := 0
	boardsLines := make([][]string, len(boards))
	for i, board := range boards {
		boardsLines[i] = strings.Split(board, "\n")
		if len(boardsLines[i]) > maxLines {
			maxLines = len(boardsLines[i])
		}
	}
	for i := 0; i < maxLines; i++ {
		for j := 0; j < len(boardsLines); j++ {
			boardLines := boardsLines[j]
			if i >= len(boardLines) {
				continue
			}
			merged += boardLines[i] + separator
		}
		merged += "\n"
	}
	return merged
}
