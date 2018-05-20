package main

import "fmt"

func main() {
	s := "asdfklasjdfliowuerijklcjjkjklsdfioj"
	sep := "jkl"
	fmt.Println(strStr(s, sep))
}
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	var i int
	for i < len(haystack) {
		j := 0
		for ; (i+j) < len(haystack) && j < len(needle) && haystack[i+j] == needle[j]; j++ {
		}
		if len(needle) != j {
			i++
		} else {
			return i
		}
	}
	return -1
}

func strStr2(haystack string, needle string) int {
	if len(needle) <= 0 {
		return 0
	}
	n := -1
	needleLen := len(needle)
	haystackLen := len(haystack)
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && i+needleLen <= haystackLen {
			if haystack[i:i+needleLen] == needle {
				n = i
				return n
			}
		}
	}
	return n
}
