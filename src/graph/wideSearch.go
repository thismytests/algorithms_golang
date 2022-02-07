package graph

import (
	"sort"
)

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func WideSearch(graph map[string][]string, start string, end string) bool {
	queue := []string{}

	queue = append(queue, start)

	for len(queue) > 0 {
		firstElemInQueue := queue[0:1][0]
		queue = queue[1:]

		if len(graph[firstElemInQueue]) == 0 {
			graph[firstElemInQueue] = []string{}
		}

		if contains(graph[firstElemInQueue], end) {
			return true
		} else {
			queue = append(queue, graph[firstElemInQueue]...)
		}

	}

	return false
}
