package main

func rotate(nums []int, k int)  {
	for ; k > len(nums); k = k - len(nums) {
	}
	tmp :=make([]int,0)
	for i := len(nums) - k; i < len(nums); i++ {
		tmp = append(tmp,nums[i])
	}
	for i := 0; i < len(nums) - k; i++ {
		tmp = append(tmp,nums[i])
	}
	for i:=0;i<len(nums) ;i++  {
		nums[i] = tmp[i]
	}
}