package linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositiveBubbleSort(t *testing.T) {
	arr := []int{3, 4, 20, 2, 1, 0, 100}
	index := LinearSearch(arr, 100)
	assert.Equal(t, index, 6)
}

func TestNegativeBubbleSort(t *testing.T) {
	arr := []int{3, 4, 20, 2, 1, 0, 100}
	index := LinearSearch(arr, 1000)
	assert.Equal(t, index, -1)
}
