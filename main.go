package main

import (
	"fmt"
	_ "net/http/pprof"
	"strings"
	"time"
)

func main() {
	maze := NewMaze(999)

	fmt.Println("Do you want to compare all algorithms or run one?")
	fmt.Println("1 - compare all")
	fmt.Println("2 - run one")
	var userChoice string
	fmt.Scan(&userChoice)

	if userChoice == "1" {
		for i := 0; i < 10; i++ {
			graph := createGraph(maze.Grid)

			dfsStart := time.Now()
			for i := 0; i < 100; i++ {
				_, err := solve("dfs", *graph)
				if err != nil {
					panic(err)
				}

			}
			dfsElapsed := time.Since(dfsStart)

			bfsStart := time.Now()
			for i := 0; i < 100; i++ {
				_, err := solve("bfs", *graph)
				if err != nil {
					panic(err)
				}

			}
			bfsElapsed := time.Since(bfsStart)

			gbfsStart := time.Now()
			for i := 0; i < 100; i++ {
				_, err := solve("gbfs", *graph)
				if err != nil {
					panic(err)
				}

			}
			gbfsElapsed := time.Since(gbfsStart)

			dijkstraStart := time.Now()
			for i := 0; i < 100; i++ {
				_, err := solve("dijkstra", *graph)
				if err != nil {
					panic(err)
				}

			}
			dijkstraElapsed := time.Since(dijkstraStart)

			fmt.Println("DFS Elapsed: ", dfsElapsed.Milliseconds())
			fmt.Println("BFS Elapsed: ", bfsElapsed.Milliseconds())
			fmt.Println("GBFS Elapsed: ", gbfsElapsed.Milliseconds())
			fmt.Println("Dijkstra Elapsed: ", dijkstraElapsed.Milliseconds())
		}
	}

	if userChoice == "2" {
		fmt.Println("Available algorithms: 'DFS', 'BFS', 'GBFS', 'dijkstra'")
		fmt.Print("Chose maze solving algorithm: ")
		var algoInput string
		fmt.Scan(&algoInput)

		graph := createGraph(maze.Grid)
		solutionPath, err := solve(strings.ToLower(algoInput), *graph)
		if err != nil {
			panic(err)
		}
		maze.CreateCoordPath(solutionPath)

		fmt.Printf("Vertices: %d\n", graph.Vertices)
		fmt.Printf("Edges: %d\n", graph.Edges)
		fmt.Printf("Start: %d\n", graph.Start.ID)
		fmt.Printf("End: %d\n", graph.End.ID)
	}
}
