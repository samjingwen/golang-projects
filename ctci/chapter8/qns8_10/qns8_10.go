package qns8_10

import (
	. "ctci/common"
)

type color struct {
	r, g, b uint8
}

type point struct {
	r, c int
}

func PaintFill(screen [][]color, pt point, newColor color) {
	R := len(screen)
	C := len(screen[0])

	// direction vectors, clockwise from north
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}

	visited := make([][]bool, R)
	for i, _ := range visited {
		visited[i] = make([]bool, C)
	}

	var q Queue[point]
	q.Enqueue(pt)
	visited[pt.r][pt.c] = true
	screen[pt.r][pt.c] = newColor

	for q.IsNotEmpty() {
		curr, _ := q.Deque()

		for i := 0; i < len(dr); i++ {
			rr := curr.r + dr[i]
			cc := curr.c + dc[i]

			if rr >= R || rr < 0 {
				continue
			}
			if cc >= C || cc < 0 {
				continue
			}
			if !visited[rr][cc] {
				q.Enqueue(point{rr, cc})
				visited[rr][cc] = true
				screen[rr][cc] = newColor
			}
		}
	}
}
