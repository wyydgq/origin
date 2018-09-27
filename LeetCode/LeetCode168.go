package main

func convertToTitle(n int) string {
	const len = 26
	var x int
	var y int
	var res string = ""
	for ;n>0 ;  {
		x = n/len
		y = n%len
		if y==0 {
			res = string('Z') +res;x--
		}else {
			res = string('A'+ y-1)+res;
		}
		n = x
	}
	return res
}