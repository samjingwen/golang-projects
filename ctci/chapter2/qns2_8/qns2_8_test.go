package qns2_8

import (
	"testing"
)
import . "ctci/chapter2"

type ArgsAndWants struct {
	node *LinkedListNode
	args *LinkedListNode
	want *LinkedListNode
}

var tests = []struct {
	name         string
	argsAndWants ArgsAndWants
}{
	{"", makeArgsAndWants1()},
	{"", makeArgsAndWants2()},
	{"", makeArgsAndWants3()},
	{"", makeArgsAndWants4()},
}

func TestLoopDetection(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoopDetection(tt.argsAndWants.args); got != tt.argsAndWants.want {
				t.Errorf("LoopDetection() = %v, want %v", got, tt.argsAndWants.want)
			}
		})
	}
}

func makeArgsAndWants1() ArgsAndWants {
	node := MakeNodeFromSlice(1)
	node.Next = node
	return ArgsAndWants{args: node, want: node}
}

func makeArgsAndWants2() ArgsAndWants {
	node := MakeNodeFromSlice(1, 2, 3, 4)
	want := node.Next
	node.Next.Next.Next.Next = want
	return ArgsAndWants{args: node, want: want}
}

func makeArgsAndWants3() ArgsAndWants {
	node := MakeNodeFromSlice(1, 2, 3)
	want := node.Next.Next
	node.Next.Next.Next = want
	return ArgsAndWants{args: node, want: want}
}

func makeArgsAndWants4() ArgsAndWants {
	node := MakeNodeFromSlice(1, 2, 3, 4, 5, 6, 7)
	want := node.Next.Next.Next.Next.Next.Next
	node.Next.Next.Next.Next.Next.Next.Next = want
	return ArgsAndWants{args: node, want: want}
}
