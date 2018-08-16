package main

import (
	"fmt"

	. "./point"
	"./queue"
)

// if the point can be add to stack
func ifAddPoint(x, y int, seen [][]bool, m [][]int, col, row int) bool {
	if x >= row || x < 0 || y >= col || y < 0 {
		return false
	}
	if m[x][y] == 1 || seen[x][y] {
		return false
	}
	return true
}

func neighPoints(x, y int) []Point {
	neighPoints := []Point{
		{X: x, Y: y - 1}, // up
		{X: x + 1, Y: y}, // right
		{X: x, Y: y + 1}, // bottom
		{X: x - 1, Y: y}, // left
	}

	return neighPoints
}

func bfs(m [][]int, col int, row int) [][]Point {
	// record the parent point of current point
	parentPointMap := [][]Point{}
	seen := [][]bool{}
	q := queue.Queue{}

	parentPointMap = make([][]Point, row)
	seen = make([][]bool, row)
	for i := 0; i < row; i++ {
		parentPointMap[i] = make([]Point, col)
		seen[i] = make([]bool, col)
	}

	q.Push(Point{0, 0})
	seen[0][0] = true

	for {
		item := q.Pop()
		if p, ok := item.(Point); ok {
			x, y := p.X, p.Y

			if x == row-1 && y == col-1 {
				break
			}

			for _, neighPoints := range neighPoints(x, y) {
				if ifAddPoint(neighPoints.X, neighPoints.Y, seen, m, col, row) {
					q.Push(neighPoints)

					nx, ny := neighPoints.X, neighPoints.Y
					seen[nx][ny] = true
					parentPointMap[nx][ny] = Point{x, y}
				}
			}
		} else {
			break
		}
	}
	fmt.Println(parentPointMap)
	return parentPointMap
}

func main() {
	m := [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	col := 5
	row := 6
	res := bfs(m, col, row)

	x, y := 5, 4
	for {
		fmt.Printf("(%d, %d)", x, y)
		if x == 0 && y == 0 {
			break
		}
		fmt.Print("<-")
		point := res[x][y]
		x, y = point.X, point.Y
	}
}
