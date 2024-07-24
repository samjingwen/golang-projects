package qns5_4

import "testing"

func TestNextBiggest(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		args args
		want uint32
	}{
		{args{0b100}, 0b1000},
		{args{0b1100}, 0b10001},
		{args{0b101100}, 0b110001},
		{args{0b101010}, 0b101100},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NextBiggest(tt.args.num); got != tt.want {
				t.Errorf("NextBiggest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextSmallest(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		args args
		want uint32
	}{
		{args{0b1000}, 0b100},
		{args{0b1100}, 0b1010},
		{args{0b101100}, 0b101010},
		{args{0b10110011}, 0b10101110},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NextSmallest(tt.args.num); got != tt.want {
				t.Errorf("NextSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
