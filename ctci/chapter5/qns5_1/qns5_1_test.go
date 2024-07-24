package qns5_1

import "testing"

func TestInsertion(t *testing.T) {
	type args struct {
		N int
		M int
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{0b10000000000, 0b10011, 2, 6}, 0b10001001100},
		{"", args{0b10000000000, 0b10011, 0, 4}, 0b10000010011},
		{"", args{0b10001111100, 0b10011, 2, 6}, 0b10001001100},
		{"", args{0b10001011100, 0b10011, 2, 6}, 0b10001001100},
		{"", args{0b10001101100, 0b10011, 2, 6}, 0b10001001100},
		{"", args{0b10001110100, 0b10011, 2, 6}, 0b10001001100},
		{"", args{0b10001110000, 0b10011, 2, 6}, 0b10001001100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insertion(tt.args.N, tt.args.M, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Insertion() = %b, want %b", got, tt.want)
			}
		})
	}
}
