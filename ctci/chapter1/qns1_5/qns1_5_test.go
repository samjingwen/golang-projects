package qns1_5

import "testing"

type args struct {
	shorter string
	longer  string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"1", args{"pale", "ale"}, true},
	{"2", args{"pale", "ple"}, true},
	{"3", args{"pales", "pale"}, true},
	{"4", args{"pale", "bale"}, true},
	{"5", args{"pake", "bake"}, true},
	{"6", args{"pale", "bake"}, false},
	{"7", args{"abcd", "abcdef"}, false},
}

func TestOneAway1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OneAway1(tt.args.shorter, tt.args.longer); got != tt.want {
				t.Errorf("OneAway1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOneAway2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OneAway2(tt.args.shorter, tt.args.longer); got != tt.want {
				t.Errorf("OneAway2() = %v, want %v", got, tt.want)
			}
		})
	}
}
