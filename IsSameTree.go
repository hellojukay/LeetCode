package main

import (
	"fmt"
)


type TreeNode struct {
      Val int
     Left *TreeNode
     Right *TreeNode
 }
 
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == q{
		return true
	}
	return p.Left == q.Left && p.Right == q.Right && isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
}

func main(){
	fmt.Println(isSameTree(nil,nil))
}