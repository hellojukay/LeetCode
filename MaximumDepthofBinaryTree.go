package LeetCode


type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 
 func maxDepth(root *TreeNode) int {
    if root == nil{
		return 0
	}
	leftLen := 1+ maxDepth(root.Left)
	rightLen := 1+ maxDepth(root.Right)
	if leftLen >rightLen{
		return leftLen
	}
	return rightLen
}