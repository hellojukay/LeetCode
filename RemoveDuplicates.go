package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var i, j = 0, 1
	for {
		if j == len(nums) {
			break
		}
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i + 1
}
func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}
