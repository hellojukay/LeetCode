package tow_sum

func twoSum(nums []int, target int) []int {
	var v1, v2 int
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if (nums[i]+nums[j] == target) && (i != j) {
				v1, v2 = i, j
				break
			}
		}
	}
	return []int{v1, v2}
}
