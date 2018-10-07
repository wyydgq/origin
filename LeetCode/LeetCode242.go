package main

import "strings"

func isAnagram(s string, t string) bool {
	m:=make(map[string]int)
	s1:=strings.Split(s,"")
	s2:=strings.Split(t,"")
	if len(s1)!=len(s2) {return false}
	for i:=0;i<len(s1);i++ {
		if _,f := m[s1[i]];f{
			m[s1[i]]++
		}else {
			m[s1[i]]=1
		}
	}
	for i:=0;i<len(s2);i++ {
		if _,f := m[s2[i]];f{
			m[s2[i]]--
		}else {
			return false
		}
	}
	for _,y:= range m{
		if y!=0{return false}
	}
	return true
}