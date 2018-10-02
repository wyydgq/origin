package main

import (
	"strings"
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	arr1 := strings.Split(s,"")
	arr2 := strings.Split(t,"")
	if len(arr1)!=len(arr2) {return false}
	if len(arr1)==0 {
		return true
	}
	m := make(map[string]string)
	for i := 0; i<len(arr1); i++ {
		if _,flag:=m[arr1[i]];!flag{
			for _,s1:=range m {
				if s1==arr2[i] {return false}
			}
			m[arr1[i]]=arr2[i]
		}
	}
	for i := 0; i<len(arr1); i++ {
		if val,flag:=m[arr1[i]];flag{
			arr1[i]=val
		}else {
			return false
		}
	}
	return strings.Join(arr1,"")==strings.Join(arr2,"")
}

func main()  {
	fmt.Println(isIsomorphic("egg","add"))
}