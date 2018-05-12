package main

import (
	"fmt"
)


func twoSum(nums []int,target int)[]int{
    var m = make(map[int]int)
    for index, value := range nums{
        m[value] = index
    }
    for index, value := range nums{
        var sub = target - value
        if v , exists := m[sub];exists{
            if v == index{
                continue
            }
            return []int{index ,v}
        }
    }
    return []int{1,1}
}


func main(){
    fmt.Printf("%+v",twoSum([]int{1,2,3,4,5},9))
}