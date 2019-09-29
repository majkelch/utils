package utils

import "strings"

// reverse string
func ReverseString(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// compare strings
func CompareStrings(a string, b string) int {
	return strings.Compare(a, b)
}