package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialState(t *testing.T) {
	assert := assert.New(t)

	g := New()

	assert.True(testBlockEmptiness("Board", g.Board, 10, t), "Board must be empty")
	assert.False(testBlockEmptiness("BlockA", g.BlockA, 5, t), "BlockA must not be empty")
	assert.False(testBlockEmptiness("BlockB", g.BlockB, 5, t), "BlockB must not be empty")
	assert.False(testBlockEmptiness("BlockC", g.BlockC, 5, t), "BlockC must not be empty")
}

func testBlockEmptiness(blockName string, block [][]BoardElement, size int, t *testing.T) bool {
	assert := assert.New(t)

	assert.Equal(size, len(block), "%s length is not equal to %d", blockName, size)

	for x := 0; x < size; x++ {
		assert.Equal(size, len(block[x]), "%s[%d] length is not equal to %d", blockName, x, size)
		for y := 0; y < size; y++ {
			if None != block[x][y] {
				return false
			}
		}
	}
	return true
}

func TestMove(t *testing.T) {
	assert := assert.New(t)

	g := New()

	g.BlockA = blockShape(1)
	g.BlockB = blockShape(1)
	g.BlockC = blockShape(1)

	if error := g.Move(A, 0, 0); error != nil {
		t.Errorf("First move of BlockA to position 0,0 must succeed. Received: %s", error)
	}
	assert.True(testBlockEmptiness("BlockA", g.BlockA, 5, t), "BlockA must be empty")

	if error := g.Move(A, 5, 5); error == nil {
		t.Errorf("Second move of empty BlockA without moving B or C to position 5,5 must fail")
	}
	g.GameOver = false
	if error := g.Move(22, 5, 5); error == nil {
		t.Errorf("Selected block (22) must fail")
	}
	g.GameOver = false

	if error := g.Move(B, 1, 0); error != nil {
		t.Errorf("First move of BlockB to position 1,0 must succeed. Received: %s", error)
	}
	assert.True(testBlockEmptiness("BlockB", g.BlockA, 5, t), "BlockA must be empty")

	if error := g.Move(C, 2, 0); error != nil {
		t.Errorf("First move of BlockC to position 2,0 must succeed. Received: %s", error)
	}

	// randomize all blocks again
	assert.False(testBlockEmptiness("BlockA", g.BlockA, 5, t), "BlockA must not be empty")
	assert.False(testBlockEmptiness("BlockB", g.BlockB, 5, t), "BlockB must not be empty")
	assert.False(testBlockEmptiness("BlockC", g.BlockC, 5, t), "BlockC must not be empty")
}

func TestMoveFullRow(t *testing.T) {
	assert := assert.New(t)

	g := New()

	g.BlockA = blockShape(8)
	g.BlockB = blockShape(8)

	if error := g.Move(A, 0, 0); error != nil {
		t.Errorf("Move of BlockA to position 0,0 must succeed")
	}

	if error := g.Move(B, 5, 0); error != nil {
		t.Errorf("Move of BlockB to position 5,0 must succeed")
	}

	assert.True(testBlockEmptiness("Board", g.Board, 10, t), "Board must be empty")
}

func TestCheckAndRemoveFullLanes(t *testing.T) {
	assert := assert.New(t)

	g := New()

	// horizontal
	for i := 0; i < 10; i++ {
		g.Board[0][i] = Green
	}

	checkAndRemoveFullLanes(g.Board)

	assert.True(testBlockEmptiness("Board", g.Board, 10, t), "Board must be empty")

	// vertical
	for i := 0; i < 10; i++ {
		g.Board[i][0] = Magenta
	}

	checkAndRemoveFullLanes(g.Board)

	assert.True(testBlockEmptiness("Board", g.Board, 10, t), "Board must be empty")

	// double
	for i := 0; i < 10; i++ {
		g.Board[i][0] = Cyan
		g.Board[0][i] = Yellow
	}

	checkAndRemoveFullLanes(g.Board)

	assert.True(testBlockEmptiness("Board", g.Board, 10, t), "Board must be empty")

	// scathered elements
	g.Board[0][0] = Cyan
	g.Board[0][1] = Yellow
	g.Board[8][8] = Blue
	g.Board[8][9] = Blue
	g.Board[9][8] = Blue
	g.Board[9][9] = Blue

	checkAndRemoveFullLanes(g.Board)

	assert.Equal(Cyan, g.Board[0][0], "Board[0][0] is not equal to Cyan")
	assert.Equal(Yellow, g.Board[0][1], "Board[0][0] is not equal to Yellow")
	assert.Equal(Blue, g.Board[8][8], "Board[8][8] is not equal to Blue")
	assert.Equal(Blue, g.Board[8][9], "Board[8][9] is not equal to Blue")
	assert.Equal(Blue, g.Board[9][8], "Board[9][8] is not equal to Blue")
	assert.Equal(Blue, g.Board[9][9], "Board[9][9] is not equal to Blue")
}

func TestGameOver(t *testing.T) {
	assert := assert.New(t)

	g := New()

	g.Board = [][]BoardElement{
		{Green, None, Green, None, Green, None, Green, None, Green, None},
		{None, Green, None, Green, None, Green, None, Green, None, Green},
		{Green, None, Green, None, Green, None, Green, None, Green, None},
		{None, Green, None, Green, None, Green, None, Green, None, Green},
		{Green, None, Green, None, Green, None, Green, None, Green, None},
		{None, Green, None, Green, None, Green, None, Green, None, Green},
		{Green, None, Green, None, Green, None, Green, None, Green, None},
		{None, Green, None, Green, None, Green, None, Green, None, Green},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None}}

	g.BlockA = blockShape(1)
	g.BlockB = blockShape(1)
	g.BlockC = blockShape(1)

	assert.False(g.isGameOver())

	g.BlockA = createContainer(5)
	g.BlockB = blockShape(10)
	g.BlockC = createContainer(5)

	assert.True(g.isGameOver())

	g.BlockA = blockShape(10)
	g.BlockB = blockShape(10)
	g.BlockC = blockShape(11)

	assert.True(g.isGameOver())

	g.BlockA = createContainer(5)
	g.BlockB = createContainer(5)
	g.BlockC = blockShape(1)

	assert.False(g.isGameOver())

	g.BlockA = blockShape(1)
	g.BlockB = blockShape(10)
	g.BlockC = createContainer(5)

	error := g.Move(A, 8, 0)

	assert.NotNil(error)

	errorGame, ok := error.(*ErrorGame)
	if !ok {
		assert.Fail("error is not of ErrorGame type")
		return
	}

	assert.Equal(GameOver, errorGame.Reason)
}

func TestPlaceBlock(t *testing.T) {
	assert := assert.New(t)

	g := New()

	// placing
	testPlacement(&g, 0, 0, blockShape(1), "blockTwoHorizontal", true, t)
	assert.Equal(Green, g.Board[0][0], "Board[0][0] is not equal to Green")
	assert.Equal(Green, g.Board[0][1], "Board[0][0] is not equal to Green")

	// out of index
	testPlacement(&g, 8, 0, blockShape(8), "blockFiveVertical", false, t)

	// placing correctly near the edge
	testPlacement(&g, 8, 8, blockShape(9), "blockTwoSquare", true, t)
	assert.Equal(Cyan, g.Board[8][8], "Board[8][8] is not equal to Cyan")
	assert.Equal(Cyan, g.Board[8][9], "Board[8][9] is not equal to Cyan")
	assert.Equal(Cyan, g.Board[9][8], "Board[9][8] is not equal to Cyan")
	assert.Equal(Cyan, g.Board[9][9], "Board[9][9] is not equal to Cyan")

	// placing over other elements
	testPlacement(&g, 9, 9, blockShape(0), "blockDot", false, t)

	// placing outside the board's range
	testPlacement(&g, -5, 5, blockShape(0), "blockDot", false, t)
	testPlacement(&g, 0, -2, blockShape(0), "blockDot", false, t)
	testPlacement(&g, -20, -2, blockShape(0), "blockDot", false, t)
	testPlacement(&g, 10, 10, blockShape(0), "blockDot", false, t)
	testPlacement(&g, 25, 0, blockShape(0), "blockDot", false, t)

	// placing near other elements with overlapping block
	testPlacement(&g, 8, 7, blockShape(0), "blockDot", true, t)
}

func testPlacement(g *Game, x int, y int, block [][]BoardElement, blockName string, shouldSucceed bool, t *testing.T) {
	if error := g.placeBlock(x, y, block); error != nil {
		if shouldSucceed {
			t.Errorf("Placement of %s at %d,%d must not return error", blockName, x, y)
		}
		return
	}
	if !shouldSucceed {
		t.Errorf("Placement of %s at %d,%d must return error", blockName, x, y)
	}
}
