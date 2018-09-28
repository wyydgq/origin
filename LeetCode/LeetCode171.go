package main

func titleToNumber(s string) int {
	var arr = string(s)
	l := len(arr)
	if l==0 {return 0}
	res :=0
	for i:=0;i<l;i++ {
		res = res*26 + int(arr[i]) - 'A'+ 1
	}
	return res
}

