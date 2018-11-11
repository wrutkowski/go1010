package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/wrutkowski/go1010/drawer"
	"github.com/wrutkowski/go1010/game"
	"github.com/wrutkowski/go1010/neural"
)

func main() {
	columns := 3
	rows := 1
	drawEveryRun := false

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
	untilTimeHasPassed := time.Now()
	loadFromFile := ""
	saveToFile := ""

	for {
		untilTimeHasPassedDiff := time.Now().Sub(untilTimeHasPassed)
		interactionEnabled := generations == 0 && steps == 0 && untilNextGeneration == false && untilFitnessIsAbove == 0 && untilTimeHasPassedDiff > 0

		if drawEveryRun || interactionEnabled {
			drawer.PrepareTerminal()

			drawer.DrawGames(columns, rows, games)

			fmt.Printf("Generation: %d\n", neuralManager.GenerationNumber())

			if interactionEnabled {
				runForSeconds := 0
				exit, generations, steps, untilNextGeneration, untilFitnessIsAbove, runForSeconds, drawEveryRun, saveToFile, loadFromFile = nextCommand(drawEveryRun)
				if runForSeconds > 0 {
					untilTimeHasPassed = time.Now().Add(time.Second * time.Duration(runForSeconds))
				}
				if saveToFile != "" {
					fmt.Print("Saving... ")
					if error := neuralManager.SaveToFile(saveToFile); error != nil {
						fmt.Println("Error while saving. ", error)
					}
					saveToFile = ""
				}
				if loadFromFile != "" {
					fmt.Print("Loading... ")
					if error := neuralManager.LoadFromFile(loadFromFile); error != nil {
						fmt.Println("Error while loading. ", error)
					}
					loadFromFile = ""
				}
			}
		}

		if exit {
			break
		}

		populationIsDead := true
		for i := 0; i < population; i++ {
			if !games[i].GameOver {
				populationIsDead = false
			}
		}
		if populationIsDead {
			if untilTimeHasPassedDiff < 0 {
				fmt.Println("Remaining running time:", untilTimeHasPassedDiff*-1)
			}
			for i := 0; i < population; i++ {
				games[i] = game.New()
			}
			neuralManager.NextGeneration()

			untilNextGeneration = false
			if generations > 0 {
				generations--
			}
		}

		for i := 0; i < population; i++ {
			output := neuralManager.Networks[i].Run(inputForGame(games[i]))
			block, x, y := outputToGameControl(output, len(games[i].Board))
			errorGame := games[i].Move(block, x, y)

			neuralManager.Networks[i].Fitness = calculateFitness(games[i], errorGame)

			if untilFitnessIsAbove > 0 && neuralManager.Networks[i].Fitness > untilFitnessIsAbove {
				untilFitnessIsAbove = 0
			}
		}

		if steps > 0 {
			steps--
		}
	}

}

func nextCommand(drawing bool) (exit bool, generations int, steps int, untilNextGeneration bool, untilFitnessIsAbove float32, runForSeconds int, drawingEnabled bool, saveToFile string, loadFromFile string) {
	instructions := `Instructions:
	Enter - next iteration
	s NUM - skip NUM of steps
	g NUM - skip NUM generations
	f NUM - until Fitness is above NUM
	t NUM - run for NUM seconds
	ng - run until next generation
	drawing on/off - enable/disable drawing each iteration
	save filename - saves top performant Neural Network to a file
	load filename - loads Neural Network to last place
	help - this help
	e - exit`

	fmt.Print("Command: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	components := strings.Split(text, " ")

	if components[0] == "exit" || components[0] == "e" {
		return true, 0, 0, false, 0, 0, drawing, "", ""
	}

	if components[0] == "ng" {
		return false, 0, 0, true, 0, 0, drawing, "", ""
	}

	if components[0] == "s" {
		if len(components) < 2 {
			fmt.Println("Wrong command format. s NUM - skip NUM of steps, eg. `s 10`")
			return nextCommand(drawing)
		}
		s, _ := strconv.Atoi(components[1])
		return false, 0, s, false, 0, 0, drawing, "", ""
	}

	if components[0] == "g" {
		if len(components) < 2 {
			fmt.Println("Wrong command format. g NUM - skip NUM generations, eg. `g 3`")
			return nextCommand(drawing)
		}
		g, _ := strconv.Atoi(components[1])
		return false, g, 0, false, 0, 0, drawing, "", ""
	}

	if components[0] == "f" {
		if len(components) < 2 {
			fmt.Println("Wrong command format. f NUM - until Fitness is above NUM, eg. `f 20`")
			return nextCommand(drawing)
		}
		f, _ := strconv.Atoi(components[1])
		return false, 0, 0, false, float32(f), 0, drawing, "", ""
	}

	if components[0] == "t" {
		if len(components) < 2 {
			fmt.Println("Wrong command format. t NUM - run for NUM seconds, eg. `t 60`")
			return nextCommand(drawing)
		}
		t, _ := strconv.Atoi(components[1])
		return false, 0, 0, false, 0, t, drawing, "", ""
	}

	if components[0] == "drawing" {
		if len(components) < 2 {
			fmt.Println("Wrong command format. drawing on/off - enable/disable drawing each iteration, eg. `drawing disable`")
			return nextCommand(drawing)
		}
		if components[1] == "enable" || components[1] == "e" || components[1] == "1" {
			return false, 0, 0, false, 0, 0, true, "", ""
		} else {
			return false, 0, 0, false, 0, 0, false, "", ""
		}
	}

	if components[0] == "save" {
		if len(components) < 2 {
			fmt.Println("Wrong command format. save filename - saves top performant Neural Network to a file, eg. `save network.neural`")
			return nextCommand(drawing)
		}
		save := components[1]
		return false, 0, 0, false, 0, 0, drawing, save, ""
	}

	if components[0] == "load" {
		if len(components) < 2 {
			fmt.Println("Wrong command format. load filename - loads Neural Network to last place, eg. `load network.neural`")
			return nextCommand(drawing)
		}
		load := components[1]
		return false, 0, 0, false, 0, 0, drawing, "", load
	}

	if components[0] == "help" {
		fmt.Println("\n" + instructions)
		return nextCommand(drawing)
	}

	return false, 0, 0, false, 0, 0, drawing, "", ""
}

func calculateFitness(g game.Game, errorGame error) float32 {
	return float32(g.Score)
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
