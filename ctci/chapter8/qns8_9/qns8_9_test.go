package qns8_9

import (
	"reflect"
	"testing"
)

func TestParens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want map[string]struct{}
	}{
		{args{1}, map[string]struct{}{"()": {}}},
		{args{2}, map[string]struct{}{"()()": {}, "(())": {}}},
		{args{3}, map[string]struct{}{"((()))": {}, "(()())": {}, "(())()": {}, "()(())": {}, "()()()": {}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Parens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParens2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want map[string]struct{}
	}{
		{args{1}, map[string]struct{}{"()": {}}},
		{args{2}, map[string]struct{}{"()()": {}, "(())": {}}},
		{args{3}, map[string]struct{}{"((()))": {}, "(()())": {}, "(())()": {}, "()(())": {}, "()()()": {}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Parens2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parens2() = %v, want %v", got, tt.want)
			}
		})
	}
}
