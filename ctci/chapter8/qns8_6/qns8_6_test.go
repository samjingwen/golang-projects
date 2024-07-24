package qns8_6

import (
	"reflect"
	"testing"
)

func TestTowersOfHanoi(t *testing.T) {
	type args struct {
		src    *Tower
		dest   *Tower
		buffer *Tower
	}
	tests := []struct {
		args args
		want *Tower
	}{
		{args{&Tower{3, 2, 1}, &Tower{}, &Tower{}}, &Tower{3, 2, 1}},
		{args{&Tower{10, 9, 8, 7, 5, 4, 3, 2, 1}, &Tower{}, &Tower{}}, &Tower{10, 9, 8, 7, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			TowersOfHanoi(tt.args.src, tt.args.dest, tt.args.buffer)
			if !tt.args.src.isEmpty() || !reflect.DeepEqual(tt.args.dest, tt.want) {
				t.Errorf("src = %v, dest = %v, buffer = %v", *tt.args.src, *tt.args.dest, *tt.args.buffer)
			}
		})
	}
}
