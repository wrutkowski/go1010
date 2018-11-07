package drawer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wrutkowski/go1010/game"

	"golang.org/x/crypto/ssh/terminal"
)

// PrepareTerminal resets caret to 0,0 position and fills entire available
// space with spaces to overwrite characters from previous draw.
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

// DrawGames draws an array of games in rows x columns grid layout
func DrawGames(columns int, rows int, games []game.Game) {
	fmt.Printf("\033[0;0H")
	for row := 0; row < rows; row++ {
		horizontalWindows := make([]string, columns)
		for column := 0; column < columns; column++ {
			index := row*columns + column
			horizontalWindows[column] = drawGame(games[index], fmt.Sprintf("Neural Network: %d", index))
		}
		fmt.Print(mergeBoardsHorizontally(" ", horizontalWindows...))
		fmt.Println()
	}
}

// DrawGame draws board and three blocks
func DrawGame(g game.Game) {
	fmt.Printf("\033[0;0H")
	fmt.Print(drawGame(g, "go1010"))
}

func drawGame(g game.Game, title string) string {
	boardDrawing := drawBoard(g.Board)
	blockADrawing := drawBoard(g.BlockA)
	blockBDrawing := drawBoard(g.BlockB)
	blockCDrawing := drawBoard(g.BlockC)

	boardDrawingLength := len(g.Board)*2 + 3
	blockDrawingLength := len(g.BlockA)*2 + 3

	// add bottom margin to blocks
	for i := 0; i < 5; i++ {
		blockADrawing += "\n"
		for j := 0; j < blockDrawingLength; j++ {
			blockADrawing += " "
		}
		blockBDrawing += "\n"
		for j := 0; j < blockDrawingLength; j++ {
			blockBDrawing += " "
		}
		blockCDrawing += "\n"
		for j := 0; j < blockDrawingLength; j++ {
			blockCDrawing += " "
		}
	}

	gameOver := ""
	if g.GameOver {
		gameOver = " - GAME OVER"
	}
	gameArea := mergeBoardsHorizontally(" ", boardDrawing, blockADrawing, blockBDrawing, blockCDrawing)
	window := windowAround(title+" | score: "+strconv.Itoa(g.Score)+gameOver, gameArea, boardDrawingLength+3*blockDrawingLength+4)

	return window
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

func windowAround(title string, content string, contentWidth int) string {
	contentLines := strings.Split(content, "\n")
	contentH := len(contentLines)
	contentW := len(title) + 5
	if contentWidth > contentW {
		contentW = contentWidth
	}

	s := "\u250C\u2500 " + title + " "
	for i := 0; i < contentW-len(title)-3; i++ {
		s += "\u2500"
	}
	s += "\u2510\n"

	for i := 0; i < contentH; i++ {
		s += "\u2502" + contentLines[i]
		for j := 0; j < contentW-contentWidth; j++ {
			s += " "
		}
		s += "\u2502\n"
	}

	s += "\u2514"
	for i := 0; i < contentW; i++ {
		s += "\u2500"
	}
	s += "\u2518\n"

	return s
}

func drawBoard(board [][]game.BoardElement) string {
	s := "  "
	for x := 0; x < len(board[0]); x++ {
		s += strconv.Itoa(x) + " "
	}
	s += " \n \u250F"
	for x := 0; x < 2*len(board[0]); x++ {
		s += "\u2501"
	}
	s += "\u2513\n"
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if y == 0 {
				s += strconv.Itoa(x) + "\u2503"
			}
			s += boardElementToString(board[x][y]) + boardElementToString(board[x][y])
			if y == len(board[x])-1 {
				s += "\u2503"
			}
		}
		s += "\n"
	}
	s += " \u2517"
	for x := 0; x < 2*len(board[0]); x++ {
		s += "\u2501"
	}
	s += "\u251B"
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
		if i < maxLines-1 {
			merged += "\n"
		}
	}
	return merged
}
