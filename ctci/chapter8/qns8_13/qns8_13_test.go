package qns8_13

import "testing"

func TestStackOfBoxes(t *testing.T) {
	type args struct {
		cuboids [][]int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{[][]int{{50, 45, 20}, {95, 37, 53}, {45, 23, 12}}}, 65},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := StackOfBoxes(tt.args.cuboids); got != tt.want {
				t.Errorf("StackOfBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
