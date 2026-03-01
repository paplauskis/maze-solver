package main

import (
	"errors"
)

type QNode struct {
	Node Node
	Next *QNode
}

type Queue struct {
	Size  int
	Start *QNode
	Tail  *QNode
}

const NoParent = -1

func BreadthFirstSearch(graph Graph, startNode, endNode Node) []Node {
	if len(graph.AdjacencyList[startNode.ID]) == 0 {
		panic(errors.New("starting vertex doesn't have any neighbors"))
	}
	graphSize := graph.End.ID * 2
	graph.Visited = make([]bool, graphSize)
	graph.Parents = make([]Node, graphSize)
	for i := range graphSize {
		graph.Parents[i] = Node{X: NoParent, Y: NoParent, ID: NoParent}
	}

	solveBFS(graph, startNode, endNode)

	return reconstructPath(graph.Parents, startNode, endNode)
}

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

func NewQueue() *Queue {
	return &Queue{
		Size:  0,
		Start: nil,
		Tail:  nil,
	}
}

func (n *Queue) Enqueue(node Node) {
	temp := &QNode{
		Node: node,
		Next: nil,
	}

	if n.Size == 0 {
		n.Start = temp
		n.Tail = temp
	} else {
		n.Tail.Next = temp
		n.Tail = temp
	}

	n.Size++
}

func (n *Queue) Dequeue() (Node, error) {
	if n.Start == nil {
		panic(errors.New("Queue is empty"))
	}

	dequeuedNode := n.Start
	n.Start = n.Start.Next
	n.Size--

	if n.Start == nil {
		n.Tail = nil
	}

	return dequeuedNode.Node, nil
}
