package main

import "fmt"

func main() {
	// nums := []int{0, 1, 0, 3, 12}
	nums := []int{0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	k := len(nums)
	for i := 0; i < k; {
		if nums[i] == 0 {
			copy(nums[i:len(nums)-1], nums[i+1:len(nums)])
			nums[len(nums)-1] = 0
			k--
		} else {
			i++
		}
	}
}
