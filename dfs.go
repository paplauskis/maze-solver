package main

import (
	"errors"
)

func DepthFirstSearch(graph Graph) []Node {
	approxGraphSize := graph.End.ID * 2
	visited := make([]bool, approxGraphSize)
	history := make([]Node, 0)
	prnts := make([]Node, graph.End.ID*2)
	for i := range prnts {
		prnts[i] = Node{ID: NoParent}
	}

	if len(graph.AdjacencyList[graph.Start.ID]) == 0 {
		panic(errors.New("starting vertex doesn't have any neighbors"))
	}

	path, _ := solveDFS(graph, graph.Start, graph.End, visited, history, prnts)

	return reconstructPath(path, graph.Start, graph.End)
}

func solveDFS(graph Graph, currNode, endNode Node, visited []bool, history, prnts []Node) ([]Node, bool) {
	if visited[currNode.ID] {
		return prnts, false
	}

	history = append(history, currNode)
	visited[currNode.ID] = true

	if currNode == endNode {
		return prnts, true
	}

	for _, node := range graph.AdjacencyList[currNode.ID] {
		var endNodeFound bool

		if visited[node.ID] {
			continue
		}
		prnts[node.ID] = currNode

		// goes to first neighbor always (depth)
		// in this case - down -> right -> left -> up (based on directions array)
		history, endNodeFound = solveDFS(graph, node, endNode, visited, history, prnts)
		if endNodeFound {
			return prnts, true
		}
	}

	return prnts, false
}
