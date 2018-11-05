package neural

import (
	"math/rand"
	"time"
)

// NetworkManager holds generation of neural networks, manages mutation and fitness
type NetworkManager struct {
	generationNumber int
	randomProvider   RandomProvider
	networks         []Network
}

// GenerationNumber returns current generation number of population of neural networks
func (manager NetworkManager) GenerationNumber() int {
	return manager.generationNumber
}

// Population returns population of neural networks
func (manager NetworkManager) Population() int {
	return len(manager.networks)
}

// NewNetworkManager returns NetworkManager configured with population of neural
// networks with neural layer: inputs + hiddenLayers + output
func NewNetworkManager(inputs int, outputs int, hiddenLayers []int, population int) NetworkManager {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomProvider := RandomProvider{randomGenerator: rand.New(randomSource)}

	var manager NetworkManager
	manager.randomProvider = randomProvider

	layers := makeLayers(inputs, outputs, hiddenLayers)

	manager.networks = make([]Network, population)
	for networkIndex := 0; networkIndex < population; networkIndex++ {
		manager.networks[networkIndex] = NewNetwork(layers, manager.randomProvider)
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
