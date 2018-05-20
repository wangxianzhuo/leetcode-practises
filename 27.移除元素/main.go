package main

import "fmt"

func main() {
	// nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	nums := []int{3, 2, 2, 3}
	// nums := []int{3, 3}
	// nums := []int{4, 5}
	l := removeElement(nums, 3)
	fmt.Println("")
	fmt.Println(l, nums)
}

func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	} else if len(nums) == 1 {
		if val == nums[0] {
			return 0
		} else {
			return 1
		}
	}

	tag := len(nums) - 1
	for i := 0; tag >= 0 && i <= tag; i++ {
		// fmt.Println("----")
		// fmt.Println("i:", i, "tag:", tag)
		// fmt.Println(nums)
		if val == nums[i] {
			for ; tag > 0 && i < tag && nums[tag] == val; tag-- {
			}
			nums[i], nums[tag] = nums[tag], nums[i]
			tag--
		}
		// fmt.Println("tag:", tag)
		// fmt.Println(nums)
		// fmt.Println("----")
	}
	return len(nums[:tag+1])
}
