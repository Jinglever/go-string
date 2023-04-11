package jgstr

import "strings"

// convert first letter to upper case
func Ucfirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}

// convert first letter to lower case
func Lcfirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[0:1]) + s[1:]
}
