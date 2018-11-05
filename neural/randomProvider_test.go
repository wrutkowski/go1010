package neural

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubRandomProvider struct {
	StubNextRange         float32
	StubNextRangeFunction func(min float32, max float32) float32
}

func (stubRandom StubRandomProvider) NextRange(min float32, max float32) float32 {
	if stubRandom.StubNextRangeFunction != nil {
		return stubRandom.StubNextRangeFunction(min, max)
	}
	return stubRandom.StubNextRange
}

func TestNextRange(t *testing.T) {
	assert := assert.New(t)

	randomSource := rand.NewSource(1) // seeded with constant so that result of pseudo random generator can be used in test
	randomProvider := RandomProvider{randomGenerator: rand.New(randomSource)}

	assert.Equal(float32(0.20932055), randomProvider.NextRange(-1, 1))
	assert.Equal(float32(0.88101816), randomProvider.NextRange(-1, 1))
	assert.Equal(float32(0.32912016), randomProvider.NextRange(-1, 1))

	assert.Equal(float32(98.87543), randomProvider.NextRange(98, 100))
	assert.Equal(float32(98.84927), randomProvider.NextRange(98, 100))
	assert.Equal(float32(99.37365), randomProvider.NextRange(98, 100))
}
