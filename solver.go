package main

import (
	"errors"
	"strings"
)

const NoParent = -1

func solve(algo string, graph Graph) ([]Node, error) {
	if len(graph.AdjacencyList[graph.Start.ID]) == 0 {
		panic(errors.New("starting vertex doesn't have any neighbors"))
	}

	// initialize arrays so no index out of bounds errors occur
	graphSize := graph.End.ID * 2
	graph.Visited = make([]bool, graphSize)
	graph.History = make([]Node, 0)
	graph.Parents = make([]Node, graphSize)
	graph.Distances = make([]int, graphSize)
	for i := range graphSize {
		graph.Parents[i] = Node{X: NoParent, Y: NoParent, ID: NoParent}
	}
	for i := range graph.Distances {
		graph.Distances[i] = 2147483648
	}

	switch strings.ToLower(algo) {
	case "dfs":
		solveDFS(graph, graph.Start, graph.End)
	case "bfs":
		solveBFS(graph, graph.Start, graph.End)
	case "gbfs":
		solveGBFS(graph, graph.Start, graph.End)
	case "dijkstra":
		graph.Distances[graph.Start.ID] = 0
		solveDijkstra(graph, graph.Start, graph.End)
	default:
		return []Node{}, errors.New("algorithm not valid")
	}

	return reconstructPath(graph.Parents, graph.Start, graph.End), nil
}
