package main

func majorityElement(nums []int) int {
	var m  = make(map[int]int)
	c := (len(nums)+1)/2
	for i := 0; i < len(nums); i++ {
		m[nums[i]] ++
		if m[nums[i]]>=c{
			return nums[i]
		}
	}
	return -1
}