package main

type ListNode struct {
    Val int
    Next *ListNode
}

//结构体函数传参不需用*取出，基本类型需要
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {return head}
	for ;head != nil; {
		if head.Val==val{
			head =head.Next
		}else {
			break
		}
	}
	if head == nil {return head}
	pre := head
	cur := head.Next
	if cur == nil {return head}
	next := head.Next.Next
	for ; cur!=nil; {
		if cur.Val == val {
			pre.Next=next
			cur=next
			if cur!=nil&&cur.Next!=nil {
				next = cur.Next
			}else {
				next = nil
			}
			continue
		}
		pre =cur
		cur=next
		if cur!=nil&&cur.Next!=nil {
			next = cur.Next
		}else {
			next = nil
		}
	}
	return head
}

func main()  {
	var l1 ListNode
	l1.Val = 1
	var l2 ListNode
	l2.Val = 2
	var l3 ListNode
	l3.Val = 3
	var l4 ListNode
	l4.Val = 4
	var l5 ListNode
	l5.Val = 5
	var l6 ListNode
	l6.Val = 6
	l1.Next=&l2
	l2.Next=&l3
	l3.Next=&l4
	l4.Next=&l5
	l5.Next=&l6
	removeElements(&l1,6)
}