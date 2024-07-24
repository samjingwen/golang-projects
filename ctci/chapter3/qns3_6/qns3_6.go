package qns3_6

import (
	"errors"
	"fmt"
)

type AnimalShelter struct {
	cats  Queue
	dogs  Queue
	order int
}

type Animal interface {
	getOrder() int
	setOrder(int)
}
type Cat struct {
	order int
}
type Dog struct {
	order int
}

func (c Cat) getOrder() int {
	return c.order
}

func (c *Cat) setOrder(order int) {
	c.order = order
}

func (d Dog) getOrder() int {
	return d.order
}

func (d *Dog) setOrder(order int) {
	d.order = order
}

func (queue *AnimalShelter) enqueue(animal Animal) {
	switch t := animal.(type) {
	case *Cat:
		animal.setOrder(queue.order)
		queue.cats.enqueue(animal)
	case *Dog:
		animal.setOrder(queue.order)
		queue.dogs.enqueue(animal)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", t, t))
	}
	queue.order++
}

func (queue *AnimalShelter) dequeAny() (Animal, error) {
	if queue.cats.isEmpty() {
		return queue.dogs.deque()
	} else if queue.dogs.isEmpty() {
		return queue.cats.deque()
	} else {
		cat, _ := queue.cats.peek()
		dog, _ := queue.dogs.peek()
		if cat.getOrder() < dog.getOrder() {
			return queue.cats.deque()
		}
		return queue.dogs.deque()
	}
}

func (queue *AnimalShelter) dequeDog() (Animal, error) {
	return queue.dogs.deque()
}

func (queue *AnimalShelter) dequeCat() (Animal, error) {
	return queue.cats.deque()
}

type Queue []Animal

func (queue *Queue) enqueue(val Animal) {
	*queue = append(*queue, val)
}

func (queue *Queue) deque() (Animal, error) {
	if len(*queue) == 0 {
		return nil, errors.New("queue is empty")
	}
	val := (*queue)[0]
	*queue = (*queue)[1:]
	return val, nil
}

func (queue *Queue) peek() (Animal, error) {
	if len(*queue) == 0 {
		return nil, errors.New("queue is empty")
	}
	return (*queue)[0], nil
}

func (queue *Queue) isEmpty() bool {
	return len(*queue) == 0
}
