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
