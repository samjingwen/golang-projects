package qns8_6

import (
	"errors"
)

func TowersOfHanoi(src, dest, buffer *Tower) {
	src.moveDisks(src.size(), dest, buffer)
}

type Tower []int

func (tower *Tower) push(val int) error {
	if tower.isEmpty() {
		*tower = append(*tower, val)
		return nil
	}
	peek, _ := tower.peek()
	if peek < val {
		return errors.New("cannot push greater value")
	}
	*tower = append(*tower, val)
	return nil
}

func (tower *Tower) pop() (int, error) {
	if !tower.isEmpty() {
		val := (*tower)[tower.size()-1]
		*tower = (*tower)[:tower.size()-1]
		return val, nil
	}
	return 0, errors.New("pop called on empty tower")
}

func (tower *Tower) peek() (int, error) {
	if !tower.isEmpty() {
		return (*tower)[tower.size()-1], nil
	}
	return 0, errors.New("peek called on empty tower")
}

func (tower *Tower) size() int {
	return len(*tower)
}

func (tower *Tower) isEmpty() bool {
	return tower.size() == 0
}

func (tower *Tower) moveTopTo(dest *Tower) {
	val, err := tower.pop()
	if err == nil {
		dest.push(val)
	}
}

func (tower *Tower) moveDisks(size int, dest *Tower, buffer *Tower) {
	if size > 0 {
		tower.moveDisks(size-1, buffer, dest)
		tower.moveTopTo(dest)
		buffer.moveDisks(size-1, dest, tower)
	}
}
