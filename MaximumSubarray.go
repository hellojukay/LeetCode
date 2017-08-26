package main

import "fmt"

// Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

// For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
// the contiguous subarray [4,-1,2,1] has the largest sum = 6.

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var maxSum = nums[0]
	var sum = nums[0]
	for index := 1; index < len(nums); index++ {
		if sum < 0 && nums[index] >= 0 {
			sum = nums[index]
			if sum > maxSum {
				maxSum = sum
			}
			continue
		}
		if sum < 0 && nums[index] < 0 {
			if sum < nums[index] {
				sum = nums[index]
				if sum > maxSum {
					maxSum = sum
				}
				continue
			}
		}
		if sum > 0 && nums[index] < 0 && (sum+nums[index]) < 0 {
			sum = 0
			if sum > maxSum {
				maxSum = sum
			}
			continue
		}
		sum = sum + nums[index]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1}))
}
