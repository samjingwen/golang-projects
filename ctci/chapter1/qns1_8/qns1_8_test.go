package qns1_8

import "testing"

type args struct {
	matrix Matrix
}

var tests = []struct {
	name string
	args args
	want [][]int
}{
	{"", args{[][]int{{1}}}, [][]int{{1}}},
	{"", args{[][]int{{1, 2, 3}}}, [][]int{{1, 2, 3}}},
	{"", args{[][]int{{1, 0, 3}}}, [][]int{{0, 0, 0}}},
	{"", args{[][]int{
		{1, 2, 3},
		{4, 0, 6}}}, [][]int{
		{1, 0, 3},
		{0, 0, 0}}},
}

func TestZeroMatrix(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := ZeroMatrix(test.args.matrix); !got.equals(test.want) {
				t.Errorf("ZeroMatrix() = %v, want %v", got, test.want)
			}
		})
	}
}
