package neural

import (
	"math"
)

// Neuron is a single cell of a neural network, containing a value
// and referencing connection to previous layer's neurons as weights
// slice
type Neuron struct {
	value   float32
	weights []float32
}

// Network conatins a 2D slice of layers of neurons
// info: neruals are not rectangular slice
type Network struct {
	Fitness float32

	neuronLayers   [][]Neuron
	randomProvider RandomProviding
}

// NewNetwork creates Network with provided input, hidden and output
// layers' configuration. Weights are randomized.
func NewNetwork(layers []int, randomProvider RandomProviding) Network {
	var network Network
	network.randomProvider = randomProvider
	network.neuronLayers = make([][]Neuron, len(layers))
	for layerIndex := 0; layerIndex < len(layers); layerIndex++ {
		network.neuronLayers[layerIndex] = make([]Neuron, layers[layerIndex])
		weightsCount := 0
		if layerIndex > 0 {
			weightsCount = layers[layerIndex-1]
		}
		for neuronIndex := 0; neuronIndex < layers[layerIndex]; neuronIndex++ {
			network.neuronLayers[layerIndex][neuronIndex] = network.newNeuron(weightsCount)
		}
	}
	return network
}

func (network Network) newNeuron(weightCount int) Neuron {
	var neuron Neuron
	if weightCount == 0 {
		return neuron
	}

	neuron.weights = make([]float32, weightCount)
	for weightIndex := 0; weightIndex < weightCount; weightIndex++ {
		neuron.weights[weightIndex] = network.randomProvider.NextRange(-1, 1)
	}
	return neuron
}

// Run method takes input values, inserts them to the first layer of
// Neural Network and runs the network using weights to previous layers'
// neurons and Tanh as activator function. Return last layer as output.
func (network Network) Run(input []float32) []float32 {
	if len(network.neuronLayers) < 3 {
		panic("Neural Network was configured incorrectly. It has less than required input, one hidden and out layer.")
	}
	if len(network.neuronLayers[0]) != len(input) {
		panic("input doesn't match first layeur of Neural Network")
	}

	// insert input values on the first layer
	for neuronIndex := 0; neuronIndex < len(network.neuronLayers[0]); neuronIndex++ {
		network.neuronLayers[0][neuronIndex].value = input[neuronIndex]
	}

	// run the network
	for layerIndex := 1; layerIndex < len(network.neuronLayers); layerIndex++ {
		for neuronInCurrentLayerIndex := 0; neuronInCurrentLayerIndex < len(network.neuronLayers[layerIndex]); neuronInCurrentLayerIndex++ {
			var value float32
			for neuronInPreviousLayerIndex := 0; neuronInPreviousLayerIndex < len(network.neuronLayers[layerIndex-1]); neuronInPreviousLayerIndex++ {
				value += network.neuronLayers[layerIndex][neuronInCurrentLayerIndex].weights[neuronInPreviousLayerIndex] * network.neuronLayers[layerIndex-1][neuronInPreviousLayerIndex].value
			}
			network.neuronLayers[layerIndex][neuronInCurrentLayerIndex].value = float32(math.Tanh(float64(value)))
		}
	}

	// return last leyers values as output
	outputLength := len(network.neuronLayers[len(network.neuronLayers)-1])
	output := make([]float32, outputLength)
	for neuronIndex := 0; neuronIndex < outputLength; neuronIndex++ {
		output[neuronIndex] = network.neuronLayers[len(network.neuronLayers)-1][neuronIndex].value
	}
	return output
}

// Mutated returns clone of the Neural Network mutating weights with small
// probability in three different ways
// - 2% to change the sign
// - 3% to add randomized 0-30%
// - 3% to substract randomized 0-30%
func (network Network) Mutated() Network {
	return network
}
