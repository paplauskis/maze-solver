package main

import "errors"

func DepthFirstSearch(graph Graph) []int {
	visited := make([]bool, graph.End*2)
	history := make([]int, 0)

	if len(graph.AdjacencyList[graph.Start]) == 0 {
		panic(errors.New("starting vertex doesn't have any neighbors"))
	}

	path, _ := solveDFS(graph, graph.Start, graph.End, visited, history)

	return path
}

func solveDFS(graph Graph, currNode, endNode int, visited []bool, history []int) ([]int, bool) {
	if visited[currNode] {
		return history, false
	}

	history = append(history, currNode)
	visited[currNode] = true

	if currNode == endNode {
		return history, true
	}

	for _, n := range graph.AdjacencyList[currNode] {
		var endNodeFound bool
		// goes to first neighbor always (depth)
		history, endNodeFound = solveDFS(graph, n, endNode, visited, history)
		if endNodeFound {
			return history, true
		}
	}

	return history, false
}
