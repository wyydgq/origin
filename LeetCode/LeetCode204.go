package main

import (
	"math"
	"fmt"
)

func countPrimes(n int) int {
	count:=0
	for i := 1; i < n; i++ {
		if isPrime(float64(i)) {count++}
	}
	return count
}

func isPrime(x float64) bool {
	if x==float64(1) {return false}
	if x==float64(2) {return true}
	max:= math.Sqrt(x)
	for i:=float64(2);i<=max ;i++  {
		if int(x)%int(i)== 0 {return false}
	}
	return true
}

func main()  {
	fmt.Println(countPrimes(10))
}