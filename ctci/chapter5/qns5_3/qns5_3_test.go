package qns5_3

import "testing"

type args struct {
	num int32
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"", args{0b1}, 2},
	{"", args{-1}, 32},
	{"", args{-11}, 30},
	{"", args{0b111}, 4},
	{"", args{0b100000}, 2},
	{"", args{0b11100000}, 4},
	{"", args{0b11011101111}, 8},
	{"", args{0b1000000000001110000000011101000}, 5},
	{"", args{-2147483648}, 2},
	{"", args{-2147483647}, 2},
}

func TestFlipBitToWin(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlipBitToWin(tt.args.num); got != tt.want {
				t.Errorf("FlipBitToWin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlipBitToWin2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlipBitToWin2(tt.args.num); got != tt.want {
				t.Errorf("FlipBitToWin2() = %v, want %v", got, tt.want)
			}
		})
	}
}
