package main

import (
	"fmt"
	"strings"
)

// todo add dijkstra
// todo compare algorithm speeds, provide more stats
func main() {
	maze := NewMaze(19)
	graph := createGraph(maze.Grid)

	var userInput string
	fmt.Println("Available algorithms: 'DFS', 'BFS', 'GBFS'")
	fmt.Print("Chose maze solving algorithm: ")
	fmt.Scan(&userInput)

	solutionPath, err := solve(strings.ToLower(userInput), *graph)
	if err != nil {
		panic(err)
	}
	maze.CreateCoordPath(solutionPath)

	fmt.Printf("Vertices: %d\n", graph.Vertices)
	fmt.Printf("Edges: %d\n", graph.Edges)
	fmt.Printf("Start: %d\n", graph.Start.ID)
	fmt.Printf("End: %d\n", graph.End.ID)
	//for i := 0; i < len(graph.AdjacencyList); i++ {
	//	fmt.Println("Node", i, "->", graph.AdjacencyList[i])
	//}
	maze.Print()
}
