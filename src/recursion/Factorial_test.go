package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, Factorial(3), 6)
}
