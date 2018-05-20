package main

import "fmt"

func main() {
	// nums := []int{1, 1}
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// nums := []int{0, 1, 2, 2, 3, 4}
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 5, 5, 5, 10, 10, 11}

	l := removeDuplicates2(nums)
	fmt.Println("")
	fmt.Println("nums", nums)
	fmt.Println("len", l)
}

func removeDuplicates(nums []int) int {
	l := 0
	i := 0
	j := i + 1
	k := len(nums)
	for ; i < k; i++ {
		l++
		for ; j < k && nums[i] == nums[j]; j++ {
		}
		if j >= k {
			// fmt.Println("j>=k, break")
			break
		}
		offset := j - (i + 1)
		if offset == 0 {
			j = i + 1
			// fmt.Println("offset == 0 continue")
			continue
		}
		copy(nums[i+1:i+1+(k-j)], nums[j:k])
		j = i + 1
		k -= offset
		// fmt.Println("i:", i, "j:", j, "k:", k, "l:", l)
		// fmt.Println("nums", nums)
	}
	return l
}

func removeDuplicates2(nums []int) int {
	var i, l int
	for i < len(nums) {
		for ; i < len(nums) && nums[i] == nums[i+1]; i++ {
		}
		if i >= len(nums) {
			l++
			break
		}
		nums[l+1] = nums[i]
		l++
	}
	return l
}

func removeDuplicates3(nums []int) int {
	cur := 1

	for i := 1; i < len(nums); i++ {
		for ; i < len(nums) && nums[i-1] == nums[i]; i++ {
		}
		if i < len(nums) {
			nums[cur] = nums[i]
			cur++
		}

	}

	return cur
}
