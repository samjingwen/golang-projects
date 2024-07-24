package common

import "testing"

func Test_elementsAreEqual(t *testing.T) {
	type args struct {
		xs []string
		ys []string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{[]string{}, []string{}}, true},
		{args{[]string{"abc"}, []string{"abc"}}, true},
		{args{[]string{"abc", "abc", "abc"}, []string{"abc", "abc", "abc"}}, true},
		{args{[]string{"abc"}, []string{"a"}}, false},
		{args{[]string{"abc", "abc", "bac"}, []string{"abc", "bac", "abc"}}, true},
		{args{[]string{"abc", "bac", "abc"}, []string{"bac", "abc"}}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ElementsAreEqual(tt.args.xs, tt.args.ys); got != tt.want {
				t.Errorf("elementsAreEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
