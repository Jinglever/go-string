package jgstr

import "strings"

// 首字母大写
func Ucfirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}
