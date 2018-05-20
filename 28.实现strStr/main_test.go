package main

import (
	"testing"
)

func Benchmark_1(b *testing.B) {
	s := "asdfklasjdfliowuerijklcjjkjklsdfioj"
	sep := "jkl"
	for i := 0; i < b.N; i++ {
		strStr(s, sep)
	}
}
func Benchmark_2(b *testing.B) {
	s := "asdfklasjdfliowuerijklcjjkjklsdfioj"
	sep := "jkl"
	for i := 0; i < b.N; i++ {
		strStr2(s, sep)
	}
}
