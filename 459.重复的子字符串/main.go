package main

import (
	"fmt"
)

func main() {
	s := "abab"
	fmt.Println(repeatedSubstringPattern(s))
}
func repeatedSubstringPattern(s string) bool {
LOOP:
	for l := 1; l <= (len(s) - l); l++ {
		sep := s[:l]
		i := l
		for (i + l) <= len(s) {
			fmt.Println("i:", i, "l:", l, sep, s[i:i+l])
			if s[i:i+l] != sep {
				continue LOOP
			}
			i += l
		}

		if len(s)-i != 0 {
			return false
		}
		return true
	}
	return false
}
func repeatedSubstringPattern2(s string) bool {
	size := len(s)
LOOP:
	for l := 1; l <= (size - l); l++ {
		if size%l != 0 {
			continue LOOP
		}
		sep := s[:l]
		for i := l; (i + l) <= size; {
			if s[i:i+l] != sep {
				continue LOOP
			}
			i = i + l
		}
		return true
	}
	return false
}
