package neural

import (
	"math/rand"
	"sort"
	"time"
)

// NetworkManager holds generation of neural networks, manages mutation and fitness
type NetworkManager struct {
	Networks []Network

	generationNumber int
	layers           []int
	randomProvider   RandomProvider
}

// GenerationNumber returns current generation number of population of neural networks
func (manager NetworkManager) GenerationNumber() int {
	return manager.generationNumber
}

// Population returns population of neural networks
func (manager NetworkManager) Population() int {
	return len(manager.Networks)
}

// NewNetworkManager returns NetworkManager configured with population of neural
// networks with neural layer: inputs + hiddenLayers + output
func NewNetworkManager(inputs int, outputs int, hiddenLayers []int, population int) NetworkManager {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomProvider := RandomProvider{randomGenerator: rand.New(randomSource)}

	var manager NetworkManager
	manager.randomProvider = randomProvider
	manager.layers = makeLayers(inputs, outputs, hiddenLayers)

	manager.Networks = make([]Network, population)
	for networkIndex := 0; networkIndex < population; networkIndex++ {
		manager.Networks[networkIndex] = NewNetwork(manager.layers, manager.randomProvider)
	}
	return manager
}

func makeLayers(inputs int, outputs int, hiddenLayers []int) []int {
	layers := make([]int, len(hiddenLayers)+2)
	layers[0] = inputs
	for index, layer := range hiddenLayers {
		layers[index+1] = layer
	}
	layers[len(hiddenLayers)+1] = outputs
	return layers
}

// NextGeneration creates new generation of Neural Networks based on
// top performing networks (fitness):
// - top performer is recreated in the new generation
// - top performer mutants are assigned to 50% slots of the new generation
// - 2nd top performer mutants are assigned to 30% slots of the new generation
// - 3rd top performer mutants are assigned to 10% slots of the new generation
// - the rest slots of the new generation are filled with randomized networks
func (manager *NetworkManager) NextGeneration() {
	manager.sortNetworksByFitness()

	population := manager.Population()
	nextGeneration := make([]Network, population)
	for networkIndex := 0; networkIndex < population; networkIndex++ {
		p := float32(networkIndex) / float32(population)
		if networkIndex == 0 {
			nextGeneration[networkIndex] = manager.Networks[0]
		} else if p < 0.5 {
			nextGeneration[networkIndex] = manager.Networks[0].Mutated()
		} else if p < 0.8 {
			nextGeneration[networkIndex] = manager.Networks[1].Mutated()
		} else if p < 0.9 {
			nextGeneration[networkIndex] = manager.Networks[2].Mutated()
		} else {
			// new Neural Network for the rest
			nextGeneration[networkIndex] = NewNetwork(manager.layers, manager.randomProvider)
		}
	}

	manager.Networks = nextGeneration
	manager.generationNumber++
}

func (manager NetworkManager) sortNetworksByFitness() {
	sort.Slice(manager.Networks, func(i, j int) bool {
		return manager.Networks[i].Fitness > manager.Networks[j].Fitness
	})
}
