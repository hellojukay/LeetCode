package main

import (
	"fmt"
)

// 这个一道垃圾题
func mySqrt(x int) int {
	var last = 0
	var mid = int(x/2) + 1
	for i := 0; i <= mid; i++ {
		r := i * i
		if r < x {
			last = i
		}
		if r == x {
			last = i
			break
		}
		if r > x {
			break
		}
	}
	return last
}

func main() {
	fmt.Println(mySqrt(0))
	fmt.Println(mySqrt(4))

}
