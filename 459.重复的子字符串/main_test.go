package main

import "testing"

// func Benchmark_1_true(b *testing.B) {
// 	s := "ababababababababababababababababa"
// 	for i := 0; i < b.N; i++ {
// 		repeatedSubstringPattern(s)
// 	}
// }
// func Benchmark_1_false(b *testing.B) {
// 	s := "abababababababababacbababababababa"
// 	for i := 0; i < b.N; i++ {
// 		repeatedSubstringPattern(s)
// 	}
// }
// func Benchmark_2_true(b *testing.B) {
// 	s := "ababababababababababababababababa"
// 	for i := 0; i < b.N; i++ {
// 		repeatedSubstringPattern2(s)
// 	}
// }
// func Benchmark_2_false(b *testing.B) {
// 	s := "abababababababababacbababababababa"
// 	for i := 0; i < b.N; i++ {
// 		repeatedSubstringPattern2(s)
// 	}
// }

func Benchmark_mode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mode()
	}
}

func mode() {
	for i := 1; i < 1000; i++ {
		_ = 1000 % i
	}
}
