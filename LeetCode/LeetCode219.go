package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i,x:= range nums{
		if _,f := m[x];f {
			if i-m[x]<=k {return true}
			m[x] = i
		}else {
			m[x] = i
		}
	}
	return false
}