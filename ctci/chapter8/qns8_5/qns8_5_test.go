package qns8_5

import "testing"

func TestRecursiveMultiply(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{0, 0}, 0},
		{args{0, 1}, 0},
		{args{3, 1}, 3},
		{args{1, 6}, 6},
		{args{3, 5}, 15},
		{args{3, 6}, 18},
		{args{5, 4}, 20},
		{args{9, 15}, 135},
		{args{88, 95}, 8360},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := RecursiveMultiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("RecursiveMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
