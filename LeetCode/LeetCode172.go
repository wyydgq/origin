package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	n = n/5
	if n<0 {return 0}
	count:=0
	for ; n>0;{
		count+=n
		n=n/5
	}
	return count
}

func main()  {
	fmt.Println(trailingZeroes(7))
}