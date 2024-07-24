package qns8_8

import (
	. "ctci/common"
	"testing"
)

func TestPermutationsWithDups(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		args args
		want []string
	}{
		{args{""}, []string{""}},
		{args{"a"}, []string{"a"}},
		{args{"aaa"}, []string{"aaa"}},
		{args{"ab"}, []string{"ab", "ba"}},
		{args{"abc"}, []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{args{"a⌘c"}, []string{"a⌘c", "ac⌘", "⌘ac", "⌘ca", "ca⌘", "c⌘a"}},
		{args{"⌘a⌘c"}, []string{"ca⌘⌘", "c⌘a⌘", "c⌘⌘a", "⌘⌘ac", "⌘⌘ca", "⌘a⌘c", "⌘ac⌘", "⌘ca⌘", "⌘c⌘a", "a⌘⌘c", "a⌘c⌘", "ac⌘⌘"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := PermutationsWithDups(tt.args.str); !ElementsAreEqual(got, tt.want) {
				t.Errorf("PermutationsWithDups() = %v, want %v", got, tt.want)
			}
		})
	}
}
