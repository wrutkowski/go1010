package game

// BoardElement represents sing object on the game board
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

// Game struct contains 10x10 game board and three 5x5 blocks of shapes
type Game struct {
	Board  [][]BoardElement
	BlockA [][]BoardElement
	BlockB [][]BoardElement
	BlockC [][]BoardElement
}

// New Game is used to initialize Game struct
func New() Game {
	g := Game{createContainer(10), createContainer(5), createContainer(5), createContainer(5)}
	return g
}

func createContainer(size int) [][]BoardElement {
	board := make([][]BoardElement, size)
	for i := 0; i < size; i++ {
		board[i] = make([]BoardElement, size)
	}
	return board
}

// BlockShape returns one of 19 shapes available in the game
func (g Game) BlockShape(number int) [][]BoardElement {
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
