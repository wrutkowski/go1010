package neural

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNetworkManager(t *testing.T) {
	assert := assert.New(t)

	manager := NewNetworkManager(2, 2, []int{3}, 10)

	assert.Equal(10, len(manager.Networks))
}

func TestMakeLayers(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{2, 3, 2}, makeLayers(2, 2, []int{3}))
	assert.Equal([]int{1, 2, 3, 4, 1}, makeLayers(1, 1, []int{2, 3, 4}))
	assert.Equal([]int{5, 10, 9, 10, 9, 10, 10}, makeLayers(5, 10, []int{10, 9, 10, 9, 10}))
}

func TestNextGeneration(t *testing.T) {
	assert := assert.New(t)

	manager := NewNetworkManager(2, 2, []int{3}, 10)
	manager.Networks[5].Fitness = 10
	network1 := manager.Networks[0]
	network2 := manager.Networks[1]
	network3 := manager.Networks[2]
	network4 := manager.Networks[3]
	network5 := manager.Networks[4]
	network6 := manager.Networks[5]
	network7 := manager.Networks[6]
	network8 := manager.Networks[7]
	network9 := manager.Networks[8]
	network10 := manager.Networks[9]

	manager.NextGeneration()

	assert.Equal(1, manager.GenerationNumber())

	assert.Equal(network6, manager.Networks[0]) // highest Fitness is moved to next generation
	assert.NotEqual(network1, manager.Networks[0])
	assert.NotEqual(network2, manager.Networks[1])
	assert.NotEqual(network3, manager.Networks[2])
	assert.NotEqual(network4, manager.Networks[3])
	assert.NotEqual(network5, manager.Networks[4])
	assert.NotEqual(network6, manager.Networks[5])
	assert.NotEqual(network7, manager.Networks[6])
	assert.NotEqual(network8, manager.Networks[7])
	assert.NotEqual(network9, manager.Networks[8])
	assert.NotEqual(network10, manager.Networks[9])
}

func TestSortNetworksByFitness(t *testing.T) {
	assert := assert.New(t)

	var manager NetworkManager
	var net1 Network
	net1.Fitness = 0.2
	var net2 Network
	net2.Fitness = 10
	var net3 Network
	net3.Fitness = 5

	manager.Networks = []Network{net1, net2, net3}

	manager.sortNetworksByFitness()

	// sort is descending
	assert.Equal(float32(10), manager.Networks[0].Fitness)
	assert.Equal(float32(5), manager.Networks[1].Fitness)
	assert.Equal(float32(0.2), manager.Networks[2].Fitness)
}
