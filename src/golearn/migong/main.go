package main

import (
	"fmt"
	"os"
)

type Point struct {
	i int
	j int
}

var directions = [4]Point{
	{-1, 0}, //上
	{1, 0}, //下
	{0, -1}, //左
	{0, 1}, //右
}
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	var row, col int
	fmt.Fscanf(file, "%d  %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

func (p Point) add(direction Point) Point {
	return Point{p.i + direction.i, p.j + direction.j}
}

func (p Point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end Point)  [][]int {
	fmt.Println(start)
	fmt.Println(end)
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	queue := []Point{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == end {
			break
		}
		for _, direction := range directions {
			next := curr.add(direction)
			val, ok := next.at(maze)
			if  ! ok || val != 0 {
				continue
			}
			val, ok = next.at(steps)
			if  ! ok || val != 0 {
				continue
			}
			if (next == start) {
				continue
			}
			currSteps, _ := curr.at(steps)
			steps[next.i][next.j] =  currSteps + 1
			queue = append(queue, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("/Users/ouyushun/work/code/goProject/golearn/migong/data")
	for _, row := range maze {
		fmt.Println(row)
	}
	steps := walk(maze, Point{0, 0}, Point{len(maze) - 1, len(maze[0]) - 1})
	fmt.Println(steps)
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
