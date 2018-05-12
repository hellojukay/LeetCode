package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) ==1{
		return len(nums)
	}
	var i = 0
	for j := 1; j < len(nums);j++{
		if nums[i] != nums[j]{
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}

func main(){
	var arr = []int{1,1,2,3,3,4,5,6,6}
	println(removeDuplicates(arr))
	fmt.Printf("%+v",arr[0:5])
}