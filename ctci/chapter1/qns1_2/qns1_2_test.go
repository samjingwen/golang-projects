package qns1_2

import "testing"

type input struct {
	str1 string
	str2 string
}

func TestCheckPermutation(t *testing.T) {
	params := map[input]bool{
		input{str1: "abc", str2: "bca"}:         true,
		input{str1: "abc", str2: "bcda"}:        false,
		input{str1: "abcz", str2: "acdg"}:       false,
		input{str1: "aabczz", str2: "azzcba"}:   true,
		input{str1: "⌘aabczz", str2: "azzcba"}:  false,
		input{str1: "aab⌘czz", str2: "azzcba"}:  false,
		input{str1: "aabczz⌘", str2: "az⌘cbza"}: true,
	}

	for param, expected := range params {
		actual := CheckPermutation(param.str1, param.str2)
		if actual != expected {
			t.Fatalf(`Expected (%v) but got Actual (%v) for inputs (%v)`,
				expected, actual, param)
		}
	}

}
