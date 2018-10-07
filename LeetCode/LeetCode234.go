package main


type ListNode struct {
    Val int
    Next *ListNode
}
 
func isPalindrome(head *ListNode) bool {
	len:=0
	tmp:=head
	for true{
		if tmp!=nil{
			len++
			tmp=tmp.Next
		}else {
			break
		}
	}
	if len==0||len==1 {return true}
	if len==2 {return head.Val==head.Next.Val}
	s:=make([]int,0)
	for i:=0;i<len;i++ {
		if i<len/2{
			s=append(s, head.Val)
		}
		if i>=(len+1)/2{
			if head.Val!=s[len-1-i]{return false}
		}
		head=head.Next
	}
	return true
}