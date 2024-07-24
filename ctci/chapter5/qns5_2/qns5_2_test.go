package qns5_2

import "testing"

func TestBinaryToString(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{0.0}, "ERROR"},
		{"", args{1.0}, "ERROR"},
		{"", args{0.72}, "ERROR"},
		{"", args{0.375}, ".011"},
		{"", args{0.015625}, ".000001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryToString(tt.args.num); got != tt.want {
				t.Errorf("BinaryToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
