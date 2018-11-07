package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wrutkowski/go1010/drawer"
	"github.com/wrutkowski/go1010/game"
	"github.com/wrutkowski/go1010/neural"
)

func main() {
	columns := 3
	rows := 4
	drawEveryRun := true

	population := rows * columns
	neuralManager := neural.NewNetworkManager(175, 3, []int{200, 230, 170, 100, 32}, population)
	games := make([]game.Game, population)

	for i := 0; i < population; i++ {
		games[i] = game.New()
	}

	exit := false
	generations := 0
	steps := 0
	untilNextGeneration := false
	untilFitnessIsAbove := float32(0)

	for {

		interactionEnabled := generations == 0 && steps == 0 && untilNextGeneration == false && untilFitnessIsAbove == 0

		if drawEveryRun || interactionEnabled {
			drawer.PrepareTerminal()

			drawer.DrawGames(columns, rows, games)

			fmt.Printf("Generation: %d\n", neuralManager.GenerationNumber())

			if interactionEnabled {
				exit, generations, steps, untilNextGeneration, untilFitnessIsAbove = nextCommand()
			}
		}

		if exit {
			break
		}

		populationIsDead := true

		for i := 0; i < population; i++ {
			output := neuralManager.Networks[i].Run(inputForGame(games[i]))
			block, x, y := outputToGameControl(output, len(games[i].Board))
			errorGame := games[i].Move(block, x, y)

			if !games[i].GameOver {
				populationIsDead = false
			}
			neuralManager.Networks[i].Fitness = calculateFitness(games[i], errorGame)

			if untilFitnessIsAbove > 0 && neuralManager.Networks[i].Fitness > untilFitnessIsAbove {
				untilFitnessIsAbove = 0
			}
		}

		if populationIsDead {
			for i := 0; i < population; i++ {
				games[i] = game.New()
			}
			neuralManager.NextGeneration()
			untilNextGeneration = false
			if generations > 0 {
				generations--
			}
		}

		if steps > 0 {
			steps--
		}
	}

}

func nextCommand() (exit bool, generations int, steps int, untilNextGeneration bool, untilFitnessIsAbove float32) {
	fmt.Print("Command [Enter - next iteration, s NUM - skip NUM of steps, g NUM - skip NUM generations, f NUM - until Fitness is above NUM, ng - run until next generation, e - exit]: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	components := strings.Split(text, " ")

	if components[0] == "exit" || components[0] == "e" {
		return true, 0, 0, false, 0
	}

	if components[0] == "ng" {
		return false, 0, 0, true, 0
	}

	if components[0] == "s" {
		if len(components) < 2 {
			fmt.Println("Wrong command format. s NUM - skip NUM of steps, eg. `s 10`")
			return nextCommand()
		}
		s, _ := strconv.Atoi(components[1])
		return false, 0, s, false, 0
	}

	if components[0] == "g" {
		if len(components) < 2 {
			fmt.Println("Wrong command format. g NUM - skip NUM generations, eg. `g 3`")
			return nextCommand()
		}
		g, _ := strconv.Atoi(components[1])
		return false, g, 0, false, 0
	}

	if components[0] == "f" {
		if len(components) < 2 {
			fmt.Println("Wrong command format. f NUM - until Fitness is above NUM, eg. `f 20`")
			return nextCommand()
		}
		f, _ := strconv.Atoi(components[1])
		return false, 0, 0, false, float32(f)
	}

	return false, 0, 0, false, 0
}

func calculateFitness(g game.Game, errorGame error) float32 {
	fitness := g.Score

	if eg, ok := errorGame.(*game.ErrorGame); ok {
		if game.IncorrectBlock == eg.Reason {
			fitness -= 10
		} else if game.IncorrectPosition == eg.Reason {
			fitness -= 10
		}
	}

	return float32(fitness)
}

func outputToGameControl(output []float32, boardSize int) (block game.BlockType, x int, y int) {
	// fmt.Println("OUTPUT:", output)
	var outputBlock game.BlockType
	if output[0] < -0.3333 {
		outputBlock = game.A
	} else if output[0] < 0.3333 {
		outputBlock = game.B
	} else {
		outputBlock = game.C
	}

	outputX := ((output[1] + 1) / 2) * float32(boardSize)
	outputY := ((output[2] + 1) / 2) * float32(boardSize)

	// fmt.Println("GAME CONTROL:", outputBlock, int(outputX), int(outputY))

	return outputBlock, int(outputX), int(outputY)
}

func inputForGame(g game.Game) []float32 {
	input := make([]float32, len(g.Board)*len(g.Board)+len(g.BlockA)*len(g.BlockA)+len(g.BlockB)*len(g.BlockB)+len(g.BlockC)*len(g.BlockC))
	i := 0
	for x := 0; x < len(g.Board); x++ {
		for y := 0; y < len(g.Board[0]); y++ {
			if g.Board[x][y] != game.None {
				input[i] = 1
			} else {
				input[i] = 0
			}
			i++
		}
	}
	for x := 0; x < len(g.BlockA); x++ {
		for y := 0; y < len(g.BlockA[0]); y++ {
			if g.BlockA[x][y] != game.None {
				input[i] = 1
			} else {
				input[i] = 0
			}
			i++
		}
	}
	for x := 0; x < len(g.BlockB); x++ {
		for y := 0; y < len(g.BlockB[0]); y++ {
			if g.BlockB[x][y] != game.None {
				input[i] = 1
			} else {
				input[i] = 0
			}
			i++
		}
	}
	for x := 0; x < len(g.BlockC); x++ {
		for y := 0; y < len(g.BlockC[0]); y++ {
			if g.BlockC[x][y] != game.None {
				input[i] = 1
			} else {
				input[i] = 0
			}
			i++
		}
	}
	// fmt.Println("INPUT:", input)
	return input
}
