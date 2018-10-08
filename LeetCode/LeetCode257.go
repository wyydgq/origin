package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
 
func binaryTreePaths(root *TreeNode) []string {
	str := make([]string,0)
	if root==nil {return nil}
	if root.Left==nil && root.Right==nil {
		str = append(str,strconv.Itoa(root.Val))
		return str
	}
	if root.Left!=nil{
		str1 := binaryTreePaths(root.Left)
		for i:=range str1{
			str1[i] = strconv.Itoa(root.Val) + "->" + str1[i]
			str = append(str,str1[i])
		}
	}
	if root.Right!=nil{
		str2 := binaryTreePaths(root.Right)
		for i:=range str2{
			str2[i] = strconv.Itoa(root.Val) + "->" + str2[i]
			str = append(str,str2[i])
		}
	}
	return str
}

func main()  {
	var t1 TreeNode
	var t2 TreeNode
	var t3 TreeNode
	var t4 TreeNode
	t1.Val=1
	t2.Val=2
	t3.Val=3
	t4.Val=4
	t1.Left = &t2
	t1.Right = &t3
	t2.Right = &t4
	fmt.Println(binaryTreePaths(nil))
}