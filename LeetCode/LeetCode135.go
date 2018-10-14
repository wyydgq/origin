package main

func candy(ratings []int) int {
	n := len(ratings)
	result := n
	increment := make([]int, n)
	inc := 1
	for i := 1;i<n;i++{
		if ratings[i] > ratings[i-1]{
			increment[i] = max(inc, increment[i])
			inc ++
		}else{
			inc = 1
		}
	}
	for i:=n-2;i>=0;i--{
		if ratings[i] > ratings[i+1]{
			increment[i] = max(inc, increment[i])
			inc++
		}else{
			inc = 1
		}
	}
	for i:=0;i<n;i++{
		result += increment[i]
	}
	return result
}
func max(a, b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}