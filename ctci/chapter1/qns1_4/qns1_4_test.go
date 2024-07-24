package qns1_4

import "testing"

type args struct {
	str string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"0", args{"Tact coa"}, true},
	{"1", args{"a"}, true},
	{"2", args{"ab"}, false},
	{"3", args{"aab"}, true},
	{"4", args{"abcdabcd"}, true},
	{"5", args{"aaaaabb"}, true},
	{"6", args{"aaaaabc"}, false},
}

func TestPalindromePermutation1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PalindromePermutation1(tt.args.str); got != tt.want {
				t.Errorf("PalindromePermutation1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPalindromePermutation2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PalindromePermutation2(tt.args.str); got != tt.want {
				t.Errorf("PalindromePermutation2() = %v, want %v", got, tt.want)
			}
		})
	}
}
