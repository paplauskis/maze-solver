package main

import "errors"

func DepthFirstSearch(graph Graph) []Node {
	visited := make([]bool, graph.End.ID*2)
	history := make([]Node, 0)

	if len(graph.AdjacencyList[graph.Start.ID]) == 0 {
		panic(errors.New("starting vertex doesn't have any neighbors"))
	}

	path, _ := solveDFS(graph, graph.Start, graph.End, visited, history)

	return path
}

// todo edit DFS to return solution path instead of history
func solveDFS(graph Graph, currNode, endNode Node, visited []bool, history []Node) ([]Node, bool) {
	if visited[currNode.ID] {
		return history, false
	}

	history = append(history, currNode)
	visited[currNode.ID] = true

	if currNode == endNode {
		return history, true
	}

	for _, node := range graph.AdjacencyList[currNode.ID] {
		var endNodeFound bool
		// goes to first neighbor always (depth)
		// in this case - down -> right -> left -> up (based on directions array)
		history, endNodeFound = solveDFS(graph, node, endNode, visited, history)
		if endNodeFound {
			return history, true
		}
	}

	return history, false
}
