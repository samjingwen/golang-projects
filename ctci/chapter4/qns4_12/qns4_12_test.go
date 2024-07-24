package qns4_12

import (
	. "ctci/chapter4"
	"testing"
)

func TestPathsWithSum(t *testing.T) {
	type args struct {
		node *TreeNode
		sum  int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{&TreeNode{Data: 2}, 2}, want: 1},

		{args: args{
			&TreeNode{
				Data: 7,
				Left: &TreeNode{Data: 3,
					Left: &TreeNode{Data: 2},
					Right: &TreeNode{Data: 6,
						Left: &TreeNode{Data: 5,
							Left: &TreeNode{Data: 4}},
						Right: &TreeNode{Data: 7}}},
				Right: &TreeNode{Data: 12,
					Left: &TreeNode{Data: 11}}}, 12}, want: 2},

		{args: args{
			&TreeNode{
				Data: 7,
				Left: &TreeNode{Data: 3,
					Left: &TreeNode{Data: 2},
					Right: &TreeNode{Data: 6,
						Left: &TreeNode{Data: 2,
							Left: &TreeNode{Data: 1}},
						Right: &TreeNode{Data: 3}}},
				Right: &TreeNode{Data: 12,
					Left: &TreeNode{Data: -5,
						Right: &TreeNode{Data: 5}}}}, 12}, want: 5},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			if got := PathsWithSum(test.args.node, test.args.sum); got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}

}
