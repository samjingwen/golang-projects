package qns3_5

import (
	"reflect"
	"testing"
)
import . "ctci/chapter3"

func TestSortStack(t *testing.T) {
	type args struct {
		stack *Stack
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"", args{NewStack(2, 1, 3)}, []interface{}{3, 2, 1}},
		{"", args{NewStack(7, 5, 1, 3)}, []interface{}{7, 5, 3, 1}},
		{"", args{NewStack(3, 5, 2, 1, 4)}, []interface{}{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if SortStack(tt.args.stack); !reflect.DeepEqual(tt.args.stack.Values(), tt.want) {
				t.Errorf("SortStack() = %v, want %v", tt.args.stack.Values(), tt.want)
			}
		})
	}
}
