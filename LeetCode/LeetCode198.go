package main

func rob(nums []int) int {
	length :=len(nums)
	if length == 0 {return 0}
	if length == 1 {return nums[0]}
	if length == 2 {
		if nums[0]>nums[1]{
			return nums[0]
		}else {
			return nums[1]
		}
	}
	if length == 3 {
		if nums[2]+nums[0]>nums[1]{
			return nums[2]+nums[0]
		}else {
			return nums[1]
		}
	}
	tmp := make([]int,len(nums))
	tmp[0]=nums[0]
	tmp[1]=nums[1]
	if nums[2]+nums[0]>nums[1]{
		tmp[2]=nums[2]+nums[0]
	}else {
		tmp[2]=nums[1]
	}
	for i := 3; i < len(nums); i++ {
		if tmp[i-2] >tmp[i-3] {
			if nums[i]+tmp[i-2]>tmp[i-1]{
				tmp[i]=nums[i]+tmp[i-2]
			}else {
				tmp[i]=tmp[i-1]
			}
		}else {
			if nums[i]+tmp[i-3]>tmp[i-1]{
				tmp[i]=nums[i]+tmp[i-3]
			}else {
				tmp[i]=tmp[i-1]
			}
		}
	}
	return tmp[length-1]
}
