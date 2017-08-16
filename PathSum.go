package main




type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

// 递归简单搞定
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil{
		return false
	}
	if root.Val == sum && root.Left== nil && root.Right == nil{
		return true
	}
	return  hasPathSum(root.Left,sum-root.Val) || hasPathSum(root.Right,sum-root.Val)
}
