package lib

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestRandomGenerator(t *testing.T) {
	r := make([]bool, 7)
	for idx := range r {
		r[idx] = false
	}

	for i := 0; i < 100; i++ {
		r[randomValueFrom(1, 6)] = true
	}

	assert.False(t, r[0])

	for i := 1; i < 7; i++ {
		assert.True(t, r[i])
	}
}

func TestRandomGeneratorA(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Printf("- %d\n", randomValueFrom(1, 6))
	}
}
