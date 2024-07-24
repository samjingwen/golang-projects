package qns4_1

import . "ctci/chapter3"

// adjacency list
type Graph [][]int

func RouteBetweenNodes(graph Graph, start, end int) bool {
	visited := make([]bool, len(graph))
	visited[start] = true

	queue := new(Queue)
	queue.Enqueue(start)

	for !queue.IsEmpty() {
		curr, _ := queue.Deque()
		node := curr.(int)
		neighbours := graph[node]
		for _, neighbour := range neighbours {
			if !visited[neighbour] {
				if neighbour == end {
					return true
				}
				queue.Enqueue(neighbour)
				visited[node] = true
			}
		}
	}
	return false
}
