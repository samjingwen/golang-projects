package qns1_7

import (
	"testing"
)

type args struct {
	image [][]int32
}

var tests = []struct {
	name string
	args args
	want [][]int32
}{
	{"", args{[][]int32{
		{1}}}, [][]int32{
		{1}}},
	{"", args{[][]int32{
		{1, 2},
		{3, 4}}}, [][]int32{
		{3, 1},
		{4, 2}}},
	{"", args{[][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}}, [][]int32{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3}}},
	{"", args{[][]int32{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}}, [][]int32{
		{13, 9, 5, 1},
		{14, 10, 6, 2},
		{15, 11, 7, 3},
		{16, 12, 8, 4}}},
	{"", args{[][]int32{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25}}}, [][]int32{
		{21, 16, 11, 6, 1},
		{22, 17, 12, 7, 2},
		{23, 18, 13, 8, 3},
		{24, 19, 14, 9, 4},
		{25, 20, 15, 10, 5}}},
}

func TestRotateMatrix(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := RotateMatrix(test.args.image); !got.equals(test.want) {
				t.Errorf("RotateMatrix() = %v, want %v", got, test.want)
			}
		})
	}
}
