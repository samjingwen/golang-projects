package qns2_6

import "testing"
import . "ctci/chapter2"

func TestPalindrome(t *testing.T) {
	type args struct {
		head *LinkedListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{MakeNodeFromSlice(1)}, true},
		{"", args{MakeNodeFromSlice(1, 2)}, false},
		{"", args{MakeNodeFromSlice(1, 1)}, true},
		{"", args{MakeNodeFromSlice(1, 2, 1)}, true},
		{"", args{MakeNodeFromSlice(1, 2, 2)}, false},
		{"", args{MakeNodeFromSlice(1, 2, 3, 4)}, false},
		{"", args{MakeNodeFromSlice(1, 2, 2, 1)}, true},
		{"", args{MakeNodeFromSlice(1, 2, 3, 2, 1)}, true},
		{"", args{MakeNodeFromSlice(1, 2, 3, 3, 2, 1)}, true},
		{"", args{MakeNodeFromSlice(1, 2, 3, 4, 5, 6)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Palindrome(tt.args.head); got != tt.want {
				t.Errorf("Palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
