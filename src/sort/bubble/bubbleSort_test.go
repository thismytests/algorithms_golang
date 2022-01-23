package bubble

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1, 0}

	BubbleSort(arr)
	fmt.Print(arr)
	assert.Equal(t, arr[0], 0)
	assert.Equal(t, arr[len(arr)-1], 5)
}
