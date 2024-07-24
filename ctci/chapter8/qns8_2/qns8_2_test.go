package qns8_2

import (
	"reflect"
	"testing"
)

func TestRobotGrid(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		args args
		want []cell
	}{
		{args{[][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}}, []cell{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}, {2, 3}}},
		{args{[][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {-1, 0, 0, 0}}}, []cell{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {2, 2}, {2, 3}}},
		{args{[][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, -1, 0}}}, []cell{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {2, 3}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := RobotGrid(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RobotGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		path []cell
	}
	tests := []struct {
		args args
		want []cell
	}{
		{args{[]cell{{0, 0}, {1, 0}, {2, 0}, {3, 0}}}, []cell{{3, 0}, {2, 0}, {1, 0}, {0, 0}}},
		{args{[]cell{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}}, []cell{{4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := reverse(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
