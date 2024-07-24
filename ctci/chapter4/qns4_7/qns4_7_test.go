package qns4_7

import (
	"reflect"
	"testing"
)

func TestBuildOrder(t *testing.T) {
	type args struct {
		projects []int
		depList  []Pair
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"", args{[]int{0, 1, 2, 3, 4, 5}, []Pair{{0, 3}, {5, 1}, {1, 3}, {5, 0}, {3, 2}}}, []int{5, 0, 1, 3, 2, 4}, false},
		{"", args{[]int{0, 1, 2, 3, 4, 5}, []Pair{{0, 3}, {5, 1}, {1, 3}, {5, 0}, {3, 2}, {0, 5}}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildOrder(tt.args.projects, tt.args.depList)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
