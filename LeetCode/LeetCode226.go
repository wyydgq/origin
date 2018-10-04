package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func invertTree(root *TreeNode) *TreeNode {
	if root==nil {return nil}
	tmp := root.Right
	root.Right = invertTree(root.Left)
	root.Left = invertTree(tmp)
	return root
}