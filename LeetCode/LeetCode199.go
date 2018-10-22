package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res:=make([]int,0)
	if root==nil {return res}
	res = append(res,root.Val)
	r1:= rightSideView(root.Left)
	r2:= rightSideView(root.Right)
	if len(r2)>=len(r1) {
		for _,v:=range r2 {
			res=append(res,v)
		}
	}else {
		for i:=0;i<len(r1);i++ {
			if i<len(r2) {
				res=append(res,r2[i])
			}else {
				res=append(res,r1[i])
			}
		}
	}
	return res
}

func main(){
	t1:= TreeNode{}
	t2:= TreeNode{}
	t3:= TreeNode{}
	t4:= TreeNode{}
	t5:= TreeNode{}
	t1.Val=1
	t2.Val=2
	t3.Val=3
	t4.Val=4
	t5.Val=5
	t1.Left=&t2
	t1.Right=&t3
	t2.Right=&t5
	t3.Right=&t4
	rightSideView(&t1)
}