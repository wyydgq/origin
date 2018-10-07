package main

func isPowerOfTwo(n int) bool {
	if n==0 {return false}
	for true{
		if n==1 {return true}
		if n%2!=0 {return false}
		n=n/2
	}
	return true
}