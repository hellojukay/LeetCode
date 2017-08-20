package main

import (
	"fmt"
)

func frogxy(n int, x, y int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return frogxy(n-x, x, y) + frogxy(n-y, x, y)
}

func main() {
	fmt.Println(frogxy(3, 1, 2))
}
