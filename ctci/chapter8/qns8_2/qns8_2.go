package qns8_2

import (
	. "ctci/common"
)

type cell struct {
	r, c int
}

// RobotGrid -1 indicates blocked cell
func RobotGrid(grid [][]int) []cell {
	R := len(grid)
	C := len(grid[0])

	visited := makeGrid[bool](R, C)
	prev := makeGrid[cell](R, C)

	// direction vectors
	dr := []int{1, 0}
	dc := []int{0, 1}

	var rq, cq Queue[int]
	rq.Enqueue(0)
	cq.Enqueue(0)
	visited[0][0] = true

	for rq.IsNotEmpty() {
		r, _ := rq.Deque()
		c, _ := cq.Deque()

		if r == R-1 && c == C-1 {
			break
		}

		for i := 0; i < len(dr); i++ {
			rr := r + dr[i]
			cc := c + dc[i]

			if rr < 0 || cc < 0 {
				continue
			}
			if rr >= R || cc >= C {
				continue
			}
			if visited[rr][cc] {
				continue
			}

			if grid[rr][cc] != -1 {
				rq.Enqueue(rr)
				cq.Enqueue(cc)

				visited[rr][cc] = true

				prev[rr][cc] = cell{r, c}
			}
		}
	}
	return reconstructPath(prev)
}

func makeGrid[T any](r, c int) [][]T {
	arr := make([][]T, r)
	for i, _ := range arr {
		arr[i] = make([]T, c)
	}
	return arr
}

func reconstructPath(prev [][]cell) []cell {
	R, C := len(prev), len(prev[0])
	start := cell{0, 0}

	var path []cell
	path = append(path, cell{R - 1, C - 1})

	at := prev[R-1][C-1]
	for at != start {
		path = append(path, at)
		at = prev[at.r][at.c]
	}
	path = append(path, start)
	return reverse(path)
}

func reverse(path []cell) []cell {
	L := len(path)
	for i := 0; i < L/2; i++ {
		path[i], path[L-i-1] = path[L-i-1], path[i]
	}
	return path
}
