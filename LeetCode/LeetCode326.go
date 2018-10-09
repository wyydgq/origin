package main

func isPowerOfThree(n int) bool {
	if n==0 {return false}
	for true {
		if n%3==0{
			n/=3
		}else {break}
	}
	if n==1 {return true}
	return false
}