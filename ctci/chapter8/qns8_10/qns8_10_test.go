package qns8_10

import "testing"

func TestPaintFill(t *testing.T) {
	type args struct {
		screen   [][]color
		pt       point
		newColor color
	}
	tests := []struct {
		args args
	}{
		{args{[][]color{{{1, 1, 1}, {1, 1, 1}}, {{1, 1, 1}, {1, 1, 1}}}, point{1, 1}, color{255, 255, 255}}},
		{args{[][]color{{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, {{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, {{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}}, point{1, 1}, color{255, 255, 255}}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			PaintFill(tt.args.screen, tt.args.pt, tt.args.newColor)
			if !allElementEqual(tt.args.screen, tt.args.newColor) {
				t.Errorf("screen = %v", tt.args.screen)
			}
		})
	}
}

func allElementEqual(screen [][]color, color color) bool {
	for _, iv := range screen {
		for _, jv := range iv {
			if jv != color {
				return false
			}
		}
	}
	return true
}
