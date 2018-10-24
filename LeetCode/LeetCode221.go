package main

import (
	"math"
	"fmt"
)



func maximalSquare(matrix [][]byte) int {
	if len(matrix)==0 {return 0}
	h := len(matrix)
	w := len(matrix[0])
	res:=make([][]byte,0)
	tmp:=make([]byte,0)
	for i:=0;i<w ;i++  {
		tmp=append(tmp,0)
	}
	for i:=0;i<h;i++{
		res=append(res,tmp)
	}
	var max byte=0
	for i:=0;i<h ;i++  {
		for j:=0;j<w ;j++  {
			if (i==0||j==0) {
				if matrix[i][j] == 1{
					res[i][j] = 1
				}else {
					res[i][j] = 0
				}
			}else {
				if matrix[i][j]==1 {
					m:= compareMin(compareMin(res[i-1][j],res[i-1][j-1]),res[i][j-1])
					res[i][j]=byte(math.Pow(float64(math.Sqrt(float64(m))+1),2))
				}else {
					res[i][j]=0
				}
			}
			max=compare(max,res[i][j])
		}
	}
	return int(max)
}

func compare(byte1,byte2 byte) byte {
	if byte1>byte2 {return byte1}
	return byte2
}

func compareMin(byte1,byte2 byte) byte {
	if byte1<byte2 {return byte1}
	return byte2
}

func main()  {
	matr:=[][]byte{{1,0,1,0,0},{1,0,1,1,1},{1,1,1,1,1},{1,0,0,1,0}}
	fmt.Println(maximalSquare(matr))
}