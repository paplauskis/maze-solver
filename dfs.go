package main

import (
	"errors"
)

func DepthFirstSearch(graph Graph) []Node {
	approxGraphSize := graph.End.ID * 2
	graph.Visited = make([]bool, approxGraphSize)
	graph.History = make([]Node, 0)
	graph.Parents = make([]Node, approxGraphSize)
	for i := range graph.Parents {
		graph.Parents[i] = Node{ID: NoParent}
	}

	if len(graph.AdjacencyList[graph.Start.ID]) == 0 {
		panic(errors.New("starting vertex doesn't have any neighbors"))
	}

	solveDFS(graph, graph.Start, graph.End)

	return reconstructPath(graph.Parents, graph.Start, graph.End)
}

func solveDFS(graph Graph, currNode, endNode Node) bool {
	if graph.Visited[currNode.ID] {
		return false
	}

	graph.History = append(graph.History, currNode)
	graph.Visited[currNode.ID] = true

	if currNode == endNode {
		return true
	}

	for _, node := range graph.AdjacencyList[currNode.ID] {
		var endNodeFound bool

		// skip visited nodes to prevent infinite cycles
		// for each new node, record its parent
		if graph.Visited[node.ID] {
			continue
		}
		graph.Parents[node.ID] = currNode

		// goes to first neighbor always (depth)
		// in this case - down -> right -> left -> up (based on directions array)
		endNodeFound = solveDFS(graph, node, endNode)
		if endNodeFound {
			return true
		}
	}

	return false
}
