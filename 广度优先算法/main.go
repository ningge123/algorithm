package main

import (
	"os"
	"fmt"
)

func main()  {
	maze := readMaze("广度优先算法/maze.in")

	fmt.Println(maze)
}

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d %d", &row, &col)

	maze := make([][]int, row)

	for i := range maze{
		maze[i] = make([]int, col)

		for j := range maze[i]{
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}