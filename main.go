package main

import (
	"time"

	"github.com/wrutkowski/go1010/drawer"
	"github.com/wrutkowski/go1010/game"
)

type boardElement int

func main() {

	g := game.New()

	drawer.PrepareTerminal()

	for i := 1; i < 19; i++ {

		g.BlockA = g.BlockShape(i)

		if i < 7 {
			g.Board[i][i] = game.BoardElement(i)
		}

		drawer.DrawGame(g)
		time.Sleep(1 * time.Second)
	}
}
