package qns2_7

import (
	"testing"
)
import . "ctci/chapter2"

type argsAndWants struct {
	node1 *LinkedListNode
	node2 *LinkedListNode
	want  *LinkedListNode
}

var tests = []struct {
	name         string
	argsAndWants argsAndWants
}{
	{"", makeArgsAndWant1()},
	{"", makeArgsAndWant2()},
	{"", makeArgsAndWant3()},
}

func TestIntersection(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.argsAndWants.node1, tt.argsAndWants.node2); got != tt.argsAndWants.want {
				t.Errorf("Intersection() = %v, want %v", got, tt.argsAndWants.want)
			}
		})
	}
}
func makeArgsAndWant1() argsAndWants {
	ptr1 := &LinkedListNode{Data: 1}
	ptr2 := ptr1
	return argsAndWants{ptr1, ptr2, ptr1}
}

func makeArgsAndWant2() argsAndWants {
	ptr1 := MakeNodeFromSlice(1, 2, 3, 4)
	ptr2 := MakeNodeFromSlice(1)
	ptr2.Next = ptr1.Next.Next.Next
	return argsAndWants{ptr1, ptr2, ptr2.Next}
}

func makeArgsAndWant3() argsAndWants {
	ptr1 := MakeNodeFromSlice(1, 2, 3, 4)
	ptr2 := MakeNodeFromSlice(1, 2)
	ptr2.Next.Next = ptr1.Next.Next.Next
	return argsAndWants{ptr1, ptr2, ptr2.Next.Next}
}
