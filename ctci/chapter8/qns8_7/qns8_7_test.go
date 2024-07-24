package qns8_7

import (
	"reflect"
	"testing"
)

func TestPermutationsWithoutDups(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		args args
		want []string
	}{
		{args{""}, []string{""}},
		{args{"a"}, []string{"a"}},
		{args{"ab"}, []string{"ab", "ba"}},
		{args{"abc"}, []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{args{"a⌘c"}, []string{"a⌘c", "ac⌘", "⌘ac", "⌘ca", "ca⌘", "c⌘a"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := PermutationsWithoutDups(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PermutationsWithoutDups() = %v, want %v", got, tt.want)
			}
		})
	}
}
