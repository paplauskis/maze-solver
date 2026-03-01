package main

type Graph struct {
	Vertices      int
	Edges         int
	AdjacencyList [][]int
	Start         int
	End           int
}

// generates graph from maze grid array
// todo try to make graph adj list more memory efficient by removing nodes that are not in crossroads
func createGraph(mazeGrid [][]int) *Graph {
	height := len(mazeGrid)
	widht := len(mazeGrid[0])

	totalNodes := height * widht
	var vertices, edges, start, end int
	adjList := make([][]int, totalNodes)

	directions := [][]int{
		{0, -1}, //up
		{1, 0},  //right
		{0, 1},  //down
		{-1, 0}, //left
	}

	for y := 0; y < height; y++ {
		for x := 0; x < widht; x++ {
			// skips walls
			if mazeGrid[y][x] == 1 {
				//(x+1 < widht && y+1 < height && mazeGrid[y+1][x] != 1 && mazeGrid[y][x+1] != 1 && y != 0) {
				continue
			}

			// generates id for node, need them to be unique because
			// nodes will be pointing to each other
			nodeID := y*widht + x
			vertices++

			//starting point (first row because maze has entrance at the top always)
			if mazeGrid[y][x] != 1 && y == 0 {
				start = nodeID
			}

			//end point (last row because ending node is always at the bottom)
			if mazeGrid[y][x] != 1 && y == height-1 {
				end = nodeID
			}

			// checks all possible neighbors
			for _, d := range directions {
				nx := x + d[0]
				ny := y + d[1]

				// checks if node is out of bounds
				if nx < 0 || ny < 0 || nx >= widht || ny >= height {
					continue
				}

				// skip walls
				if mazeGrid[ny][nx] == 1 {
					continue
				}

				neighborID := ny*widht + nx
				edges++
				// add node to adjacency list
				adjList[nodeID] = append(adjList[nodeID], neighborID)
			}
		}
	}

	return &Graph{
		Vertices:      vertices,
		Edges:         edges,
		AdjacencyList: adjList,
		Start:         start,
		End:           end,
	}
}
