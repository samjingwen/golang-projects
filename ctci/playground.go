package main

import (
	"container/list"
	"fmt"
)

type color struct {
	r, g, b uint8
}

func main() {
	l := list.New()

	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushFront(4)

	fmt.Println(l.Front())

	l.Remove(l.Front())

	fmt.Println(l.Front())
	fmt.Println(l.Back())

	l.Remove(l.Back())

	fmt.Println(l.Front())
	fmt.Println(l.Back())

}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

func (h *IntHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}
