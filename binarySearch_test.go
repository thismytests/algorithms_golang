package golang_algoritms

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	//arr := []int{0, 3, 2, 5, 6, 8, 23, 9, 4, 2, 1, 2, 9, 6, 4, 1, 7, -1, -5, 23, 6, 2, 35, 6, 3, 32, 9, 4, 2, 1, 2, 9, 6, 4, 1, 7, -1, -5, 23, 9, 4, 2, 1, 2, 9, 6, 4, 1, 7, -1, -5, 23}
	arr := []int{5, 4, 3, 2, 1, 0}

	BinarySearch(arr)
	fmt.Print(arr)
}
