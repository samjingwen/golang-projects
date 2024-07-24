package qns8_4

import (
	"reflect"
	"testing"
)

func TestPowerSet(t *testing.T) {
	type args struct {
		set []int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{args{[]int{}}, [][]int{{}}},
		{args{[]int{1}}, [][]int{{}, {1}}},
		{args{[]int{1, 2}}, [][]int{{}, {2}, {1}, {2, 1}}},
		{args{[]int{1, 2, 3}}, [][]int{{}, {3}, {2}, {3, 2}, {1}, {3, 1}, {2, 1}, {3, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := PowerSet(tt.args.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PowerSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
