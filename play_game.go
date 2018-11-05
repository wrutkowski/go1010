package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wrutkowski/go1010/drawer"
	"github.com/wrutkowski/go1010/game"
)

type boardElement int

func main() {

	g := game.New()

	drawer.PrepareTerminal()
	drawer.DrawGame(g)

	for {
		block, x, y, exit := nextMoveInteractive()
		if exit {
			break
		}

		gameOver, _ := g.Move(block, x, y)

		drawer.PrepareTerminal()
		drawer.DrawGame(g)

		if gameOver {
			fmt.Println("GAME OVER")
			break
		}
	}
}

func nextMoveInteractive() (block game.BlockType, x int, y int, exit bool) {
	fmt.Print("Next move: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	if text == "exit" || text == "e" {
		return 0, 0, 0, true
	}

	components := strings.Split(text, " ")
	if len(components) != 3 {
		fmt.Println("Provide 3 components: 0-2 for block number and X and Y position of block placement...")
		return nextMoveInteractive()
	}

	blockNumber, _ := strconv.Atoi(components[0])
	positionX, _ := strconv.Atoi(components[1])
	positionY, _ := strconv.Atoi(components[2])

	var selectedBlock game.BlockType
	switch blockNumber {
	case 0:
		selectedBlock = game.A
	case 1:
		selectedBlock = game.B
	case 2:
		selectedBlock = game.C
	}

	return selectedBlock, positionX, positionY, false
}
