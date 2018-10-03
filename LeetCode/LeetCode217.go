package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _,x:= range nums{
		if _,f := m[x];f {
			return true
		}else {
			m[x] = 1
		}
	}
	return false
}
