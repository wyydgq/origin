package main

func isHappy(n int) bool {
	collection  := make([]int,0)
	for ; ; {
		collection = append(collection, n)
		sum :=0
		for ;n>0; {
			x := n%10
			n=n/10
			sum += x*x
		}
		if sum==1 {return true}
		for i:=0;i< len(collection);i++ {
			if collection[i]==sum {return false}
		}
		n = sum
	}
}