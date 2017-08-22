package main

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[0] == target {
		return 0
	}
	if nums[0] > target {
		return 0
	}
	return 1 + searchInsert(nums[1:], target)
}
