package qns1_9

import "testing"

type args struct {
	s1, s2 string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"", args{"waterbottle", "erbottlewat"}, true},
	{"", args{"abc", "bac"}, false},
	{"", args{"abcd", "dabc"}, true},
}

func TestStringRotation(t *testing.T) {
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			if got := StringRotation(test.args.s1, test.args.s2); got != test.want {
				t.Errorf("StringRotation() = %v, want %v", got, test.want)
			}
		})
	}
}
