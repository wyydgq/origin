package main

import (
	"strings"
	"fmt"
	"strconv"
)

func compareVersion(version1 string, version2 string) int {
	arr1 := strings.Split(version1,".")
	arr2 := strings.Split(version2,".")
	var lenth int
	if len(arr1)>len(arr2){
		lenth = len(arr2)
	}else {
		lenth = len(arr1)
	}
	for i := 0; i < lenth; i++ {
		if compareVer(arr1[i],arr2[i])>0 {
			return 1
		}else if compareVer(arr1[i],arr2[i])<0 {
			return -1
		}
	}
	if len(arr1)==len(arr2) {return 0}
	if len(arr1)>len(arr2){
		for i:=lenth;i<len(arr1);i++{
			t,_:=strconv.ParseInt(arr1[i],10,64)
			if t !=0 {return 1}
		}
		return 0
	}else {
		for i:=lenth;i<len(arr2);i++{
			t,_:=strconv.ParseInt(arr2[i],10,64)
			if t !=0 {return -1}
		}
		return 0
	}
	return compareVer(strconv.Itoa(len(arr1)),strconv.Itoa(len(arr2)))
}

func compareVer(s1 string, s2 string) int {
	i1,_ := strconv.ParseInt(s1,10,64)
	i2,_ := strconv.ParseInt(s2,10,64)
	if i1 > i2 {return 1}
	if i1 < i2 {return -1}
	return 0
}

func main()  {
	fmt.Println(compareVersion("1.0","1"))
}