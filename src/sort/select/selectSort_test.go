package _select

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{100, 4, 200, 2, 0}

	SelectSort(arr)
	fmt.Print(arr)
	assert.Equal(t, 0, arr[0])
	assert.Equal(t, arr[len(arr)-1], 200)
}
