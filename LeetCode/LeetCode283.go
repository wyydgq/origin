package main

func moveZeroes(nums []int)  {
	s:=len(nums)
	if s==0 {return }
	c:=0
	idx:=0
	for _,x:=range nums{
		if x==0 {c++}
	}
	for _,x:=range nums{
		if x!=0 {
			nums[idx]=x
			idx++
		}
	}
	for i:=s-c;i<s;i++ {
		nums[i]=0
	}
}