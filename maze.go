package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Maze struct {
	Size int
	Grid [][]int
}

func NewMaze(size int) *Maze {
	if size%2 == 0 {
		size++
	}

	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
		for j := range grid[i] {
			grid[i][j] = 1 // wall
		}
	}

	m := &Maze{
		Size: size,
		Grid: grid,
	}

	rand.Seed(time.Now().UnixNano())

	m.generate(1, 1)

	startX := randomOdd(1, size-2)
	m.Grid[1][startX] = 0
	m.Grid[0][startX] = 2

	endX := randomOdd(1, size-2)
	m.Grid[size-2][endX] = 0
	m.Grid[size-1][endX] = 3

	return m
}

func randomOdd(min, max int) int {
	for {
		n := rand.Intn(max-min+1) + min
		if n%2 == 1 {
			return n
		}
	}
}

func (m *Maze) generate(x, y int) {
	m.Grid[y][x] = 0

	directions := [][]int{
		{0, -2},
		{2, 0},
		{0, 2},
		{-2, 0},
	}

	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	for _, d := range directions {
		nx := x + d[0]
		ny := y + d[1]

		if nx > 0 && nx < m.Size-1 &&
			ny > 0 && ny < m.Size-1 &&
			m.Grid[ny][nx] == 1 {

			m.Grid[y+d[1]/2][x+d[0]/2] = 0
			m.generate(nx, ny)
		}
	}
}

// changes maze grid cells of solution path to another value
func (maze *Maze) CreateCoordPath(nodePath []Node) {
	for _, node := range nodePath {
		maze.Grid[node.Y][node.X] = 7
	}
}

func (m *Maze) Print() {
	for _, row := range m.Grid {
		for _, cell := range row {
			switch cell {
			case 0:
				fmt.Print("   ")
			case 1:
				fmt.Print("███")
			case 7:
				fmt.Print(" ◯ ")
			}
		}
		fmt.Println()
	}
}
