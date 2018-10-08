package main

func missingNumber(nums []int) int {
	max := (len(nums)+1)*len(nums)/2
	for _,x:=range nums{
		max -= x
	}
	return max
}