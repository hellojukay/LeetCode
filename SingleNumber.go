package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var n = nums[0]
	for i := 1; i < len(nums); i++{
		n = n ^ nums[i]
	}
	return n
}

func main(){
	fmt.Println(singleNumber([]int{11,22,33,11,39,22,33}))
}