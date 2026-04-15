package main

import (
	"errors"
)

func solveBFS(graph Graph, startNode, endNode Node) {
	queue := NewQueue()
	graph.Visited[startNode.ID] = true
	queue.Enqueue(startNode)

	for {
		if queue.Size == 0 {
			return
		}
		node, _ := queue.Dequeue()
		neighbors := graph.AdjacencyList[node.ID]

		for _, x := range neighbors {
			if graph.Visited[x.ID] {
				continue
			}

			queue.Enqueue(x)
			graph.Visited[x.ID] = true
			graph.Parents[x.ID] = node
			graph.History = append(graph.History, x)

			if x == endNode {
				return
			}
		}
	}
}
