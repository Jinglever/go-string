package jgstr

import (
	"strconv"
	"strings"
	"unicode"
)

// convert CamelCase to camel_case
func CamelToSnake(s string) string {
	var buf strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			buf.WriteRune('_')
		}
		buf.WriteRune(unicode.ToLower(r))
	}
	return buf.String()
}

func FloatVal(s string) float64 {
	n, err := strconv.ParseFloat(s, 64) //nolint
	if err != nil {
		n = 0
	}
	return n
}
