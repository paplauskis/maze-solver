package main

import "errors"

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

func BreadthFirstSearch(graph Graph, startNode, endNode Node) ([]Node, []Node) {

	if len(graph.AdjacencyList[startNode.ID]) == 0 {
		panic(errors.New("starting vertex doesn't have any neighbors"))
	}

	path, history := solveBFS(graph, startNode, endNode)

	return reconstructPath(path, startNode, endNode), history
}

func solveBFS(graph Graph, startNode, endNode Node) ([]Node, []Node) {
	graphSize := graph.End.ID * 2
	visited := make([]bool, graphSize)

	//stores parents, f.e. if vertex 0 is the starting point and it goes to vertex 3,
	//vertex 3's parent is 0, if vertex 3 has a child vertex 5 (for example),
	//then vertex 5's parent is 3 and so on....
	prnts := make([]Node, graphSize)
	var history []Node

	for i := range graphSize {
		prnts[i] = Node{X: NoParent, Y: NoParent, ID: NoParent}
	}

	queue := NewQueue()
	visited[startNode.ID] = true
	queue.Enqueue(startNode)

	for {
		if queue.Size == 0 {
			return prnts, history
		}
		node, _ := queue.Dequeue()
		neighbors := graph.AdjacencyList[node.ID]

		for _, x := range neighbors {
			if visited[x.ID] {
				continue
			}

			queue.Enqueue(x)
			visited[x.ID] = true
			prnts[x.ID] = node
			history = append(history, x)

			if x == endNode {
				return prnts, history
			}
		}
	}
}

// reconstructs path from vertex's parents (latest search)
//
//	0  1  2  3  4  5  (array indexes)
//
// [-1, 0, 0, 1, 0, 3] (array values, -1 is root vertex)
// if path needs to be found from vertex 0 to vertex 3, then returned array
// will be [0, 1, 3] as a graph it will look like 0 -> 1 -> 3
func reconstructPath(prnts []Node, startNode, endNode Node) []Node {
	var path []Node
	for i := endNode; i.ID != NoParent; i = prnts[i.ID] {
		path = append(path, i)
	}

	reversedPath := reverse(path) //reverse for readability
	if reversedPath[0] == startNode {
		return reversedPath
	}

	return []Node{}
}

func reverse(arr []Node) []Node {
	newArr := make([]Node, len(arr))

	for i := 0; i < len(arr); i++ {
		newArr[i] = arr[len(arr)-1-i]
	}

	return newArr
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
