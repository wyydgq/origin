package main

import "strings"

func wordPattern(pattern string, str string) bool {
	m:=make(map[string]string)
	m2:=make(map[string]string)
	arr1:=strings.Split(pattern,"")
	arr2:=strings.Split(str," ")
	if len(arr1)!=len(arr2){return false}
	for i := 0; i < len(arr1); i++ {
		if v,f:=m[arr1[i]];f{
			if v!=arr2[i] {return false}
		}else {
			if _,f:=m2[arr2[i]];f{return false}
			m[arr1[i]] = arr2[i]
			m2[arr2[i]] = arr1[i]
		}
	}
	return true
}