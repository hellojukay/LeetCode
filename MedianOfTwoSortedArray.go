package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var length = len(nums1) + len(nums2)
	temp := []int{}
	var i = 0
	for {
		if i == length {
			break
		}
		i = i + 1
		if len(nums1) > 0 && len(nums2) > 0 {
			if nums1[0] < nums2[0] {
				temp = append(temp, nums1[0])
				nums1 = nums1[1:]
			} else {
				temp = append(temp, nums2[0])
				nums2 = nums2[1:]
			}
		}
		if len(nums1) > 0 && len(nums2) == 0 {
			temp = append(temp, nums1[0])
			nums1 = nums1[1:]
		}
		if len(nums1) == 0 && len(nums2) > 0 {
			temp = append(temp, nums2[0])
			nums2 = nums2[1:]
		}
	}
	fmt.Println(temp)
	if length%2 == 0 {
		return (float64(temp[length/2-1]) + float64(temp[length/2])) / 2
	}
	return float64(temp[length/2])
}

func main() {
	var a = []int{2, 3}
	var b = []int{}
	print(findMedianSortedArrays(a, b))
}
