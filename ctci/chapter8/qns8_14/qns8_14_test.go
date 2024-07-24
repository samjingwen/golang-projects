package qns8_14

import "testing"

func Test_countEval(t *testing.T) {
	type args struct {
		expr   string
		result bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"1^0|0|1", false}, 2},
		{"1", args{"0&0&0&1^1|0", true}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countEval(tt.args.expr, tt.args.result); got != tt.want {
				t.Errorf("countEval() = %v, want %v", got, tt.want)
			}
		})
	}

}
