package neural

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"strconv"
	"strings"
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
	fmt.Printf("Network[0] fitness: %.0f        \n", manager.Networks[0].Fitness)
	fmt.Printf("Network[1] fitness: %.0f        \n", manager.Networks[1].Fitness)
	fmt.Printf("Network[2] fitness: %.0f        \n", manager.Networks[2].Fitness)
	fmt.Print("\033[4A")
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

// SaveToFile sorts neural network by their fitness and saves the top performant
// network to file with a given name, returns an error in case of a save failure
func (manager NetworkManager) SaveToFile(name string) error {
	neuralSerialized := ""

	manager.sortNetworksByFitness()

	network := manager.Networks[0]

	for layer := 0; layer < len(manager.layers); layer++ {
		neuralSerialized += strconv.Itoa(manager.layers[layer])
		if layer < len(manager.layers)-1 {
			neuralSerialized += ","
		}
	}

	neuralSerialized += "|"

	for layerIndex := 0; layerIndex < len(network.neuronLayers); layerIndex++ {
		for neuronIndex := 0; neuronIndex < len(network.neuronLayers[layerIndex]); neuronIndex++ {
			for weightIndex := 0; weightIndex < len(network.neuronLayers[layerIndex][neuronIndex].weights); weightIndex++ {
				neuralSerialized += fmt.Sprintf("%f,", network.neuronLayers[layerIndex][neuronIndex].weights[weightIndex])
			}
		}
	}
	neuralSerialized = neuralSerialized[:len(neuralSerialized)-1]

	err := ioutil.WriteFile(name, []byte(neuralSerialized), 0644)

	return err
}

// LoadFromFile loads and parses content of the file with given name and replaces
// least performant network with parsed one, returns an error in case the load or
// parse was not successful
func (manager NetworkManager) LoadFromFile(name string) error {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}

	content := string(data)

	contentComponents := strings.Split(content, "|")
	if len(contentComponents) != 2 {
		return fmt.Errorf("Incorrect file format")
	}

	// layers
	layersComponents := strings.Split(contentComponents[0], ",")

	if len(layersComponents) != len(manager.layers) {
		return fmt.Errorf("Incompatible layer setting")
	}
	for layerIndex := 0; layerIndex < len(manager.layers); layerIndex++ {
		parsedLayer, parseError := strconv.Atoi(layersComponents[layerIndex])
		if parseError != nil {
			return parseError
		}
		if parsedLayer != manager.layers[layerIndex] {
			return fmt.Errorf("Incompatible layer setting. Parsed: %d, expected: %d", parsedLayer, manager.layers[layerIndex])
		}
	}

	// weights
	weightsComponents := strings.Split(contentComponents[1], ",")
	loadedWeightIndex := 0
	network := manager.Networks[0]
	var loadedNetwork Network
	loadedNetwork.randomProvider = network.randomProvider
	loadedNetwork.neuronLayers = make([][]Neuron, len(network.neuronLayers))
	for layerIndex := 0; layerIndex < len(network.neuronLayers); layerIndex++ {
		loadedNetwork.neuronLayers[layerIndex] = make([]Neuron, len(network.neuronLayers[layerIndex]))
		for neuronIndex := 0; neuronIndex < len(network.neuronLayers[layerIndex]); neuronIndex++ {
			var neuron Neuron
			neuron.weights = make([]float32, len(network.neuronLayers[layerIndex][neuronIndex].weights))
			for weightIndex := 0; weightIndex < len(network.neuronLayers[layerIndex][neuronIndex].weights); weightIndex++ {
				if loadedWeightIndex > len(weightsComponents)-1 {
					return fmt.Errorf("Incompatible weights length")
				}
				parsedWeight, parseError := strconv.ParseFloat(weightsComponents[loadedWeightIndex], 32)
				if parseError != nil {
					return parseError
				}
				neuron.weights[weightIndex] = float32(parsedWeight)

				loadedWeightIndex++
			}
			loadedNetwork.neuronLayers[layerIndex][neuronIndex] = neuron
		}
	}
	manager.Networks[len(manager.Networks)-1] = loadedNetwork

	return nil
}
