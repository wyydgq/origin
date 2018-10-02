package main

func reverseList(head *ListNode) *ListNode {
	if head==nil {return nil}
	if head.Next==nil {return head}
	res := reverseList(head.Next)
	head.Next.Next=head
	head.Next=nil
	return res
}

func main()  {
	var l1 ListNode
	var l2 ListNode
	var l3 ListNode
	l1.Val=1
	l2.Val=2
	l3.Val=3
	l1.Next=&l2
	l2.Next=&l3
	reverseList(&l1)
}
