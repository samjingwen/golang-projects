package qns1_6

import "testing"

func TestStringCompression(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"a"}, "a"},
		{"", args{"aa"}, "a2"},
		{"", args{"aaa"}, "a3"},
		{"", args{"abc"}, "abc"},
		{"", args{"aab"}, "aab"},
		{"", args{"aabc"}, "aabc"},
		{"", args{"aaabbcc"}, "a3b2c2"},
		{"", args{"aabcccccaaa"}, "a2b1c5a3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringCompression(tt.args.str); got != tt.want {
				t.Errorf("StringCompression(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
