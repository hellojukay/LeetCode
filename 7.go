package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x >= 0 && x < 10 {
		return x
	}
	var xx = x
	if x < 0 {
		xx = -x
	}
	var s string
	s = fmt.Sprintf("%d", xx)
	k := 1
	y := 0
	for i := 0; i < len(s); i++ {
		y = y + (int(s[i])-int('0'))*k
		k = k * 10
	}
	if x < 0 {
		y = -y
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		y = 0
	}
	return y
}

func main() {
	var x = reverse(-2147483648)
	fmt.Println(x)
}
