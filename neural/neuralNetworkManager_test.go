package neural

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNetworkManager(t *testing.T) {
	assert := assert.New(t)

	// stubRandomProvider := StubRandomProvider{StubNextRange: 0.25}

	manager := NewNetworkManager(2, 2, []int{3}, 10)

	assert.Equal(10, len(manager.networks))
}

func TestMakeLayers(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{2, 3, 2}, makeLayers(2, 2, []int{3}))
	assert.Equal([]int{1, 2, 3, 4, 1}, makeLayers(1, 1, []int{2, 3, 4}))
	assert.Equal([]int{5, 10, 9, 10, 9, 10, 10}, makeLayers(5, 10, []int{10, 9, 10, 9, 10}))
}
