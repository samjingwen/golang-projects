package qns8_3

import "testing"

func TestMagicIndex(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{[]int{0, 1, 3, 4, 5}}, 1},
		{args{[]int{1, 2, 3, 4, 5}}, -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MagicIndex(tt.args.arr); got != tt.want {
				t.Errorf("MagicIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagicIndexUnique(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{[]int{1, 2, 3, 4, 4}}, 4},
		{args{[]int{1, 2, 3, 4, 4}}, 4},
		{args{[]int{-10, -5, 2, 2, 2, 3, 4, 7, 9, 12, 13}}, 7},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MagicIndexUnique(tt.args.arr); got != tt.want {
				t.Errorf("MagicIndexUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
