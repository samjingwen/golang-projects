package qns8_1

import "testing"

func TestTripleStep(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{4}, 7},
		{args{5}, 13},
		{args{10}, 274},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := TripleStep(tt.args.n); got != tt.want {
				t.Errorf("TripleStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
