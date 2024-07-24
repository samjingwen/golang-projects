package qns8_12

import (
	"testing"
)

func TestNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		size int
	}{
		{args{1}, 1},
		{args{2}, 0},
		{args{3}, 0},
		{args{4}, 2},
		{args{8}, 92},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NQueens(tt.args.n); len(got) != tt.size {
				t.Errorf("NQueens() = %v, want %v", got, tt.size)
			}
		})
	}
}
