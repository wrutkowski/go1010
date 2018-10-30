package drawer

import (
	"fmt"
	"strings"

	"github.com/wrutkowski/go1010/game"

	"golang.org/x/crypto/ssh/terminal"
)

func PrepareTerminal() {
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

func DrawGame(g game.Game) {
	boardDrawing := drawBoard(g.Board)
	blockADrawing := drawBoard(g.BlockA)
	blockBDrawing := drawBoard(g.BlockB)
	blockCDrawing := drawBoard(g.BlockC)

	display := mergeBoardsHorizontally(" ", boardDrawing, blockADrawing, blockBDrawing, blockCDrawing)

	fmt.Printf("\033[0;0H")
	fmt.Print(display)
}

func boardElementToString(element game.BoardElement) string {
	switch element {
	case game.None:
		return "\033[40m \033[0m"
	case game.Red:
		return "\033[41m \033[0m"
	case game.Green:
		return "\033[42m \033[0m"
	case game.Yellow:
		return "\033[43m \033[0m"
	case game.Blue:
		return "\033[44m \033[0m"
	case game.Magenta:
		return "\033[45m \033[0m"
	case game.Cyan:
		return "\033[46m \033[0m"
	case game.White:
		return "\033[47m \033[0m"
	}
	return ""
}

func drawBoard(board [][]game.BoardElement) string {
	s := "\u250F"
	for x := 0; x < 2*len(board[0]); x++ {
		s += "\u2501"
	}
	s += "\u2513\n"
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if y == 0 {
				s += "\u2503"
			}
			s += boardElementToString(board[x][y]) + boardElementToString(board[x][y])
			if y == len(board[x])-1 {
				s += "\u2503"
			}
		}
		s += "\n"
	}
	s += "\u2517"
	for x := 0; x < 2*len(board[0]); x++ {
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
