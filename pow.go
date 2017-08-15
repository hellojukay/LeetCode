package main

import (
	"fmt"
)

// 效率低，测试不通过
// func myPow(x float64, n int) float64 {
// 	if n < 0 {
// 		x = 1 / x
// 		n = -n
// 	}
// 	if n == 0 {
// 		return 1
// 	}
// 	return x * myPow(x, n-1)
// }

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}
	return x * myPow(x, n-1)
}

func main() {
	fmt.Print(myPow(2.0, 3))
}
