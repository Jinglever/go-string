package jgstr

import (
	"sort"
	"strings"
)

// CombineSortedWithSep combines a slice of strings into a single string,
// with the given separator between each string. The strings are sorted.
func CombineSortedWithSep(sep string, strs []string) string {
	sort.Strings(strs)
	return strings.Join(strs, sep)
}
