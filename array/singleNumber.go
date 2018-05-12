package main

func singleNumber(nums []int) int {
	var i = nums[0]
	for index := 1 ;index < len(nums); index++{
		i ^= nums[index]
	}
	return i
}

func main(){
	var arr = []int{2,2,1,2,2}
	println(singleNumber(arr))
}