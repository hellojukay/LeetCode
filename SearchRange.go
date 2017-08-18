package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	var i, j = -1, -1
	var status = true
	for index := range nums {
		if nums[index] == target {
			if status {
				i = index
				status = false
			}
			j = index
		} else {
			if !status {
				break
			}
		}
	}
	return []int{i, j}
}

func main() {
	fmt.Println(searchRange([]int{1, 2}, 3))
}
