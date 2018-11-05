package neural

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	stubRandomProvider := StubRandomProvider{StubNextRange: 0.25}
	neuralNetwork := NewNetwork([]int{2, 3, 2}, stubRandomProvider)

	assert.Equal(3, len(neuralNetwork.neuronLayers))

	assert.Equal(2, len(neuralNetwork.neuronLayers[0]))
	assert.Equal(3, len(neuralNetwork.neuronLayers[1]))
	assert.Equal(2, len(neuralNetwork.neuronLayers[2]))

	assert.Equal(0, len(neuralNetwork.neuronLayers[0][0].weights))
	assert.Equal(0, len(neuralNetwork.neuronLayers[0][1].weights))
	assert.Equal(2, len(neuralNetwork.neuronLayers[1][0].weights))
	assert.Equal(2, len(neuralNetwork.neuronLayers[1][1].weights))
	assert.Equal(2, len(neuralNetwork.neuronLayers[1][2].weights))
	assert.Equal(3, len(neuralNetwork.neuronLayers[2][0].weights))
	assert.Equal(3, len(neuralNetwork.neuronLayers[2][1].weights))

	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[1][0].weights[0])
	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[1][0].weights[1])
	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[1][1].weights[0])
	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[1][1].weights[1])
	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[1][2].weights[0])
	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[1][2].weights[1])
	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[2][0].weights[0])
	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[2][0].weights[1])
	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[2][0].weights[2])
	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[2][1].weights[0])
	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[2][1].weights[1])
	assert.Equal(float32(0.25), neuralNetwork.neuronLayers[2][1].weights[2])
}

func TestRun(t *testing.T) {
	assert := assert.New(t)

	stubRandomProvider := StubRandomProvider{StubNextRange: 0.8}
	neuralNetwork := NewNetwork([]int{2, 3, 2}, stubRandomProvider)

	input := []float32{-0.4, 0.5}
	output := neuralNetwork.Run(input)

	assert.Equal(float32(-0.4), neuralNetwork.neuronLayers[0][0].value)
	assert.Equal(float32(0.5), neuralNetwork.neuronLayers[0][1].value)

	assert.Equal([]float32{0.18928106, 0.18928106}, output)
}

func TestMutated(t *testing.T) {
	assert := assert.New(t)

	var stubRandomProvider StubRandomProvider
	stubRandomProvider.StubNextRangeFunction = func(min float32, max float32) float32 {
		if min == -1 && max == 1 {
			return float32(0.25)
		}
		if min == 0 && max == 100 {
			return float32(1.5)
		}

		if min == 0 && max == 0.3 {
			return float32(0.2)
		}
		return float32(0)
	}
	neuralNetwork := NewNetwork([]int{2, 3, 2}, stubRandomProvider)

	mutant := neuralNetwork.Mutated()

	assert.Equal(3, len(mutant.neuronLayers))

	assert.Equal(2, len(mutant.neuronLayers[0]))
	assert.Equal(3, len(mutant.neuronLayers[1]))
	assert.Equal(2, len(mutant.neuronLayers[2]))

	assert.Equal(0, len(mutant.neuronLayers[0][0].weights))
	assert.Equal(0, len(mutant.neuronLayers[0][1].weights))
	assert.Equal(2, len(mutant.neuronLayers[1][0].weights))
	assert.Equal(2, len(mutant.neuronLayers[1][1].weights))
	assert.Equal(2, len(mutant.neuronLayers[1][2].weights))
	assert.Equal(3, len(mutant.neuronLayers[2][0].weights))
	assert.Equal(3, len(mutant.neuronLayers[2][1].weights))

	assert.Equal(float32(0.2), mutant.neuronLayers[1][0].weights[0])
	assert.Equal(float32(0.25), mutant.neuronLayers[1][0].weights[1])
	assert.Equal(float32(0.25), mutant.neuronLayers[1][1].weights[0])
	assert.Equal(float32(0.25), mutant.neuronLayers[1][1].weights[1])
	assert.Equal(float32(0.25), mutant.neuronLayers[1][2].weights[0])
	assert.Equal(float32(0.25), mutant.neuronLayers[1][2].weights[1])
	assert.Equal(float32(0.25), mutant.neuronLayers[2][0].weights[0])
	assert.Equal(float32(0.25), mutant.neuronLayers[2][0].weights[1])
	assert.Equal(float32(0.25), mutant.neuronLayers[2][0].weights[2])
	assert.Equal(float32(0.25), mutant.neuronLayers[2][1].weights[0])
	assert.Equal(float32(0.25), mutant.neuronLayers[2][1].weights[1])
	assert.Equal(float32(0.25), mutant.neuronLayers[2][1].weights[2])
}
