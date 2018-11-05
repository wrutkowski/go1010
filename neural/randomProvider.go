package neural

import (
	"math/rand"
	"time"
)

// RandomProviding is an interface for pseud random generator
type RandomProviding interface {
	NextRange(min float32, max float32) float32
}

// RandomProvider implements RandomProviding
type RandomProvider struct {
	randomGenerator *rand.Rand
}

// NewRandomProvider return RandomProvider configured with
// math/rand generator and seeding with current time in nanoseconds
func NewRandomProvider() RandomProvider {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomProvider := RandomProvider{randomGenerator: rand.New(randomSource)}
	return randomProvider
}

// NextRange return pseudo random number between min and max (exclusive)
func (random RandomProvider) NextRange(min float32, max float32) float32 {
	if max < min {
		panic("min must be less than max")
	}
	return min + ((max - min) * random.randomGenerator.Float32())
}
