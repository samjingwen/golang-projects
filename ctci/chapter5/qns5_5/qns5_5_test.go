package qns5_5

import "testing"

func TestConversion(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{0b0, 0b0}, 0},
		{args{0b1, 0b1}, 0},
		{args{0b1, 0b0}, 1},
		{args{0b11101, 0b01111}, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Conversion(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Conversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
