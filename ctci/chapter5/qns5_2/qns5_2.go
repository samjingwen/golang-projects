package qns5_2

import (
	"fmt"
	"math"
	"strings"
)

func BinaryToString(num float64) string {
	if num <= 0 || num >= 1 {
		return "ERROR"
	}

	var builder strings.Builder
	_, _ = fmt.Fprint(&builder, ".")

	size := 0
	for size < 32 && num > 0 {
		num *= 2
		if num >= 1 {
			_, _ = fmt.Fprint(&builder, "1")
			num -= 1
		} else {
			_, _ = fmt.Fprint(&builder, "0")
		}

		num = math.Round(num*1000000) / 1000000
		size++
	}

	if size >= 32 {
		return "ERROR"
	}
	return builder.String()
}
