package qns1_3

import (
	"testing"
)

type args struct {
	url        []rune
	expected   []rune
	trueLength int
}

func TestURLify(t *testing.T) {
	tests := []struct {
		name string
		args args
	}{
		{"1", makeArgs("Mr John Smith", "Mr%20John%20Smith")},
		{"2", makeArgs("12 ⌘3", "12%20⌘3")},
		{"2", makeArgs("abc123", "abc123")},
		{"2", makeArgs(" abc123", "%20abc123")},
		{"2", makeArgs(" abc123 ", "%20abc123%20")},
	}
	for _, tt := range tests {
		got := &tt.args.url
		t.Run(tt.name, func(t *testing.T) {
			URLify(&tt.args.url, tt.args.trueLength)
			if string(*got) != string(tt.args.expected) {
				t.Errorf("URLify() = %s, want %s", string(*got), string(tt.args.expected))
			}
		})
	}
}

func makeArgs(orig string, expected string) args {
	origRunes := []rune(orig)
	expectedRunes := []rune(expected)

	inputRunes := make([]rune, len(origRunes), len(expectedRunes))
	copy(inputRunes, origRunes)
	inputRunes = inputRunes[:len(origRunes)]

	return args{inputRunes, expectedRunes, len(origRunes)}
}
