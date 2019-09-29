package utils

import "testing"

func TestReverseString(t *testing.T) {
	res := ReverseString("Hello")
	exp := "olleH"
	if res != exp {
		t.Errorf("Expected: %s, got: %s", exp, res)
	}
}

func TestCompareStrings(t *testing.T) {
	res := CompareStrings("first string", "second string")
	exp := -1
	if res != exp {
		t.Errorf("Expected %v, got: %v", exp, res)
	}
}
