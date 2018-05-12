package main

func containsDuplicate(nums []int) bool {
    if len(nums) < 2 {
		return false
	}
	var m = make(map[int]bool)
	for _, value := range nums{
		if exsits , _ := m[value];exsits{
			return true
		}
		m[value]=true
	}
	return false
}

func main(){
	println(containsDuplicate([]int{1,1}))
	println(containsDuplicate([]int{1,2,3}))

}