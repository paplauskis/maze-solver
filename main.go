package main

import "fmt"

func main() {
	maze := NewMaze(7)
	maze.Print()

	graph := createGraph(maze.Grid)

	for i := 0; i < len(graph.AdjacencyList); i++ {
		fmt.Println("Node", i, "→", graph.AdjacencyList[i])
	}

	fmt.Printf("Vertices: %d\n", graph.Vertices)
	fmt.Printf("Edges: %d\n", graph.Edges)
	fmt.Printf("Start: %d\n", graph.Start)
	fmt.Printf("End: %d\n", graph.End)

	path := DepthFirstSearch(*graph)

	fmt.Println(path)
}
