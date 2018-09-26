package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	var res = make([]int,2)
	for x:=0;x< len(numbers);x++{
		for y:=x+1;y< len(numbers);y++{
			if numbers[x]+numbers[y]==target{
				res[0] = x+1
				res[1] = y+1
			}
		}
	}
	return res
}

func main()  {
	s:=[]int{2,7,11,15}
	fmt.Println(twoSum(s,9))
}