package main

import "testing"

func Benchmark_1(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 5, 5, 5, 10, 10, 11}
	for i := 0; i < b.N; i++ {
		removeDuplicates(nums)
	}
}

func Benchmark_2(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 5, 5, 5, 10, 10, 11}
	for i := 0; i < b.N; i++ {
		removeDuplicates2(nums)
	}
}
func Benchmark_3(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 5, 5, 5, 10, 10, 11}
	for i := 0; i < b.N; i++ {
		removeDuplicates3(nums)
	}
}
