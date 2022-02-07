package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWideSearchNegative(t *testing.T) {
	graph := make(map[string][]string)
	graph["a"] = []string{"b", "c"}

	assert.Equal(t, false, WideSearch(graph, "a", "d"))
}

func TestWideSearchPositive1(t *testing.T) {
	graph := make(map[string][]string)
	graph["a"] = []string{"b", "c"}
	graph["b"] = []string{}
	graph["c"] = []string{}

	assert.Equal(t, true, WideSearch(graph, "a", "b"))
}

func TestWideSearchPositive2(t *testing.T) {
	graph := make(map[string][]string)
	graph["a"] = []string{"b", "c"}
	graph["b"] = []string{"e"}
	assert.Equal(t, true, WideSearch(graph, "a", "e"))
}
