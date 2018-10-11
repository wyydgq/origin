package main

func canCompleteCircuit(gas []int, cost []int) int {
	res:=make([]int,0)
	for i,_:=range gas{
		res=append(res, gas[i]-cost[i])
	}
	s:=0
	for _,v:=range res{
		s+=v
	}
	if s<0 {return -1}
	s2:=0
	r:=0
	for i,v:=range res{
		if v+s2>=0 {
			s2+=v
		}else {
			s2=0
			r=i+1
		}
	}
	return r
}