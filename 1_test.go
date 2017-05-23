package tow_sum

import (
	"testing"
)

func Test_towSum(t *testing.T) {
	var nums = twoSum([]int{3, 3}, 6)
	if nums[0] != 0 || nums[1] != 1 {
		t.Error("failed")
	}
}
