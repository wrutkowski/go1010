package game

import (
	"fmt"
	"math/rand"
)

// BoardElement represents single object on the game board
type BoardElement int

// BoardElement can be one of the following colors
const (
	None    BoardElement = 0
	Red     BoardElement = 1
	Green   BoardElement = 2
	Yellow  BoardElement = 3
	Blue    BoardElement = 4
	Magenta BoardElement = 5
	Cyan    BoardElement = 6
	White   BoardElement = 7
)

// BlockType represents one of the three blocks available in the game
type BlockType int

// BlockType can be one of the following values
const (
	A BlockType = 0
	B BlockType = 1
	C BlockType = 2
)

// ErrorGameReason represents numeric reason passed with the GameError struct
type ErrorGameReason int

// ErrorGameReasons
const (
	GameOver          ErrorGameReason = 1
	IncorrectBlock    ErrorGameReason = 2
	IncorrectPosition ErrorGameReason = 3
)

// ErrorGame struct implements Error interface and is used to communicate
// about errors in the game
type ErrorGame struct {
	Reason  ErrorGameReason
	Message string
}

func (e *ErrorGame) Error() string {
	return e.Message
}

// Game struct contains 10x10 game board and three 5x5 blocks of shapes
type Game struct {
	Board    [][]BoardElement
	BlockA   [][]BoardElement
	BlockB   [][]BoardElement
	BlockC   [][]BoardElement
	Score    int
	GameOver bool

	randomGenerator *rand.Rand
}

// New Game is used to initialize Game struct: 10x10 board and three randomly
// assigned blocks
func New() Game {
	var g Game
	// randomSource := rand.NewSource(time.Now().UnixNano())
	randomSource := rand.NewSource(1)
	g.randomGenerator = rand.New(randomSource)
	g.Board = createContainer(10)
	g.assignRandomBlocks()
	return g
}

// Move is used to select one of available blocks and place it on x,y position.
// In case the placement is not possible or block does not exist error is returned.
func (g *Game) Move(block BlockType, x int, y int) error {
	if g.GameOver {
		return &ErrorGame{GameOver, "Cannot continue playing game in a game over state"}
	}

	var selectedBlock [][]BoardElement
	switch block {
	case A:
		selectedBlock = g.BlockA
	case B:
		selectedBlock = g.BlockB
	case C:
		selectedBlock = g.BlockC
	default:
		g.GameOver = true
		return &ErrorGame{IncorrectBlock, fmt.Sprintf("Incorrect block type specified (%d)", block)}
	}

	if isBlockEmpty(selectedBlock) {
		g.GameOver = true
		return &ErrorGame{IncorrectBlock, "Selected block is empty"}
	}

	error := g.placeBlock(x, y, selectedBlock)
	if error != nil {
		g.GameOver = true
		return error
	}

	emptyBlock := createContainer(5)
	switch block {
	case A:
		g.BlockA = emptyBlock
	case B:
		g.BlockB = emptyBlock
	case C:
		g.BlockC = emptyBlock
	}

	if isBlockEmpty(g.BlockA) && isBlockEmpty(g.BlockB) && isBlockEmpty(g.BlockC) {
		g.assignRandomBlocks()
	}

	if g.isGameOver() {
		g.GameOver = true
		return &ErrorGame{GameOver, "No other move is possible. Game over."}
	}

	return nil
}

func (g *Game) assignRandomBlocks() {
	g.BlockA = g.randomShape()
	g.BlockB = g.randomShape()
	g.BlockC = g.randomShape()
}

// placeBlock places provided block on the board at x and y position being 0,0 block's position.
// In case the placement is not possible error is returned.
func (g *Game) placeBlock(x int, y int, block [][]BoardElement) error {
	if x < 0 || y < 0 {
		return &ErrorGame{IncorrectPosition, fmt.Sprintf("%d,%d is below 0,0", x, y)}
	}

	newBoard := make([][]BoardElement, len(g.Board))
	for i := range g.Board {
		newBoard[i] = make([]BoardElement, len(g.Board[i]))
		copy(newBoard[i], g.Board[i])
	}

	placementScore := 0

	for boardX := x; boardX < x+len(block); boardX++ {
		for boardY := y; boardY < y+len(block); boardY++ {
			blockX := boardX - x
			blockY := boardY - y
			if boardX >= len(newBoard) || boardY >= len(newBoard) {
				if block[blockX][blockY] != None {
					return &ErrorGame{IncorrectPosition, fmt.Sprintf("%d,%d is out of board and block at %d,%d is not empty (%d)", boardX, boardY, blockX, blockY, block[blockX][blockY])}
				}
				continue
			}
			if newBoard[boardX][boardY] == None {
				newBoard[boardX][boardY] = block[blockX][blockY]
				if block[blockX][blockY] != None {
					placementScore++
				}
			} else if block[blockX][blockY] != None {
				return &ErrorGame{IncorrectPosition, fmt.Sprintf("board at %d,%d is not empty (%d) and block at %d,%d is also not empty (%d)", boardX, boardY, newBoard[boardX][boardY], blockX, blockY, block[blockX][blockY])}
			}
		}
	}

	fullLanesScore := checkAndRemoveFullLanes(newBoard)

	g.Board = newBoard
	g.Score += placementScore + fullLanesScore
	return nil
}

func (g *Game) isMovePossible(block [][]BoardElement, x int, y int) bool {
	for boardX := x; boardX < x+len(block); boardX++ {
		for boardY := y; boardY < y+len(block); boardY++ {
			blockX := boardX - x
			blockY := boardY - y
			if boardX >= len(g.Board) || boardY >= len(g.Board) {
				if block[blockX][blockY] != None {
					return false
				}
				continue
			}
			if g.Board[boardX][boardY] != None && block[blockX][blockY] != None {
				return false
			}
		}
	}

	return true
}

// checkAndRemoveFullLanes firstly counts all full rows and columns
// and then removes them from the board, replacing with None value
func checkAndRemoveFullLanes(board [][]BoardElement) int {
	fullRows := make([]bool, len(board))
	fullCols := make([]bool, len(board))
	for i := 0; i < len(board); i++ {
		fullRows[i] = true
		fullCols[i] = true
	}

	score := 0

	// check all full rows and columns before removing anything
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board); y++ {
			if fullRows[x] {
				fullRows[x] = board[x][y] != None
			}
			if fullCols[y] {
				fullCols[y] = board[x][y] != None
			}
		}
	}

	// remove rows
	for x := 0; x < len(fullRows); x++ {
		if !fullRows[x] {
			continue
		}

		score += len(fullCols)

		for y := 0; y < len(fullCols); y++ {
			board[x][y] = None
		}
	}

	// remove columns
	for y := 0; y < len(fullCols); y++ {
		if !fullCols[y] {
			continue
		}

		score += len(fullRows)

		for x := 0; x < len(fullRows); x++ {
			board[x][y] = None
		}
	}

	return score
}

func (g *Game) isGameOver() bool {
	movePossible := false
	if !isBlockEmpty(g.BlockA) && g.isMovePossibleAnywhere(g.BlockA) {
		movePossible = true
	}
	if !isBlockEmpty(g.BlockB) && g.isMovePossibleAnywhere(g.BlockB) {
		movePossible = true
	}
	if !isBlockEmpty(g.BlockC) && g.isMovePossibleAnywhere(g.BlockC) {
		movePossible = true
	}

	return !movePossible
}

func (g *Game) isMovePossibleAnywhere(block [][]BoardElement) bool {
	for x := 0; x < len(g.Board); x++ {
		for y := 0; y < len(g.Board[x]); y++ {
			if g.Board[x][y] != None {
				continue
			}

			if g.isMovePossible(block, x, y) {
				return true
			}
		}
	}

	return false
}

// randomShape returns random shape from BlockShape method
func (g *Game) randomShape() [][]BoardElement {
	return blockShape(g.randomGenerator.Intn(19))
}

// blockShape returns one of 19 shapes available in the game
func blockShape(number int) [][]BoardElement {
	switch number {
	case 0:
		return [][]BoardElement{
			{Red, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 1:
		return [][]BoardElement{
			{Green, Green, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 2:
		return [][]BoardElement{
			{Green, None, None, None, None},
			{Green, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 3:
		return [][]BoardElement{
			{Yellow, Yellow, Yellow, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 4:
		return [][]BoardElement{
			{Yellow, None, None, None, None},
			{Yellow, None, None, None, None},
			{Yellow, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 5:
		return [][]BoardElement{
			{Blue, Blue, Blue, Blue, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 6:
		return [][]BoardElement{
			{Blue, None, None, None, None},
			{Blue, None, None, None, None},
			{Blue, None, None, None, None},
			{Blue, None, None, None, None},
			{None, None, None, None, None}}
	case 7:
		return [][]BoardElement{
			{Magenta, Magenta, Magenta, Magenta, Magenta},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 8:
		return [][]BoardElement{
			{Magenta, None, None, None, None},
			{Magenta, None, None, None, None},
			{Magenta, None, None, None, None},
			{Magenta, None, None, None, None},
			{Magenta, None, None, None, None}}
	case 9:
		return [][]BoardElement{
			{Cyan, Cyan, None, None, None},
			{Cyan, Cyan, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 10:
		return [][]BoardElement{
			{White, White, White, None, None},
			{White, White, White, None, None},
			{White, White, White, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 11:
		return [][]BoardElement{
			{Cyan, Cyan, Cyan, None, None},
			{Cyan, None, None, None, None},
			{Cyan, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 12:
		return [][]BoardElement{
			{Cyan, Cyan, Cyan, None, None},
			{None, None, Cyan, None, None},
			{None, None, Cyan, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 13:
		return [][]BoardElement{
			{None, None, Cyan, None, None},
			{None, None, Cyan, None, None},
			{Cyan, Cyan, Cyan, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 14:
		return [][]BoardElement{
			{Cyan, None, None, None, None},
			{Cyan, None, None, None, None},
			{Cyan, Cyan, Cyan, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 15:
		return [][]BoardElement{
			{White, White, None, None, None},
			{White, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 16:
		return [][]BoardElement{
			{White, White, None, None, None},
			{None, White, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 17:
		return [][]BoardElement{
			{None, White, None, None, None},
			{White, White, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	case 18:
		return [][]BoardElement{
			{White, None, None, None, None},
			{White, White, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}

	default:
		return [][]BoardElement{
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None},
			{None, None, None, None, None}}
	}
}

// createContainer returns square two diemnsional slice of size x size
func createContainer(size int) [][]BoardElement {
	board := make([][]BoardElement, size)
	for i := 0; i < size; i++ {
		board[i] = make([]BoardElement, size)
	}
	return board
}

// isBlockEmpty checks if all fields of the block are None
func isBlockEmpty(block [][]BoardElement) bool {
	for x := 0; x < len(block); x++ {
		for y := 0; y < len(block); y++ {
			if block[x][y] != None {
				return false
			}
		}
	}
	return true
}
