package qns8_11

import "testing"

func TestCoins(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{1}, 1},
		{args{2}, 1},
		{args{5}, 2},
		{args{10}, 4}, // 10 ones; 5 ones 1 five; 2 fives; 1 tens;
		{args{25}, 13},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Coins(tt.args.n); got != tt.want {
				t.Errorf("Coins() = %v, want %v", got, tt.want)
			}
		})
	}
}
