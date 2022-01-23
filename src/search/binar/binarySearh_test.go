package binar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositiveBinarySearch(t *testing.T) {
	searchedItem := 5
	arr := []int{0, 1, 2, 3, 4, searchedItem, 6, 7, 8, 9}
	index := BinarySearch(arr, searchedItem)
	assert.Equal(t, searchedItem, index)
}

func TestPositiveBinarySearchRight(t *testing.T) {
	searchedItem := 3
	arr := []int{0, 1, 2, searchedItem}
	index := BinarySearch(arr, searchedItem)
	assert.Equal(t, searchedItem, index)
}

func TestPositiveBinarySearchLeft(t *testing.T) {
	searchedItem := 3
	arr := []int{searchedItem, 1, 2, 3}
	index := BinarySearch(arr, searchedItem)
	assert.Equal(t, searchedItem, index)
}

func TestNegativeBinarySearch(t *testing.T) {
	notExistItem := 100
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := BinarySearch(arr, notExistItem)
	assert.Equal(t, index, -1)
}
