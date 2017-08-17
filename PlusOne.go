package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	var length = len(digits)
	if digits[length-1] < 9 {
		digits[length-1]++
		return digits
	}
	return append(plusOne(digits[0:length-1]), 0)
}

func main() {
	fmt.Println(plusOne([]int{1, 9}))
}
