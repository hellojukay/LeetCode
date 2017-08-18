package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	var length = len(nums)
	if length == 0 {
		return 0
	}
	var i, j = 0, length - 1
	for {
		if i > j {
			break
		}
		//挖坑
		if nums[i] == val {
			for {
				if nums[j] != val {
					// 埋坑
					nums[i], nums[j] = nums[j], nums[i]
					j--
					break
				} else {
					j--
					if i > j {
						break
					}
				}
			}
		} else {
			i++
		}

	}
	fmt.Println(nums)
	return i
}

func main() {
	fmt.Println(removeElement([]int{4, 4, 4}, 4))
}
