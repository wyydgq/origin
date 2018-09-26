package main

import "fmt"

type MinStack struct{
	s []int
	min int
	size int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	var m MinStack
	m.size=0
	return m
}


func (this *MinStack) Push(x int)  {
	if this.size==0{
		this.s=append(this.s, x)
		this.size=1
		this.min=0
	}else {
		this.s=append(this.s, x)
		if x<this.s[this.min]{
			this.min=this.size
		}
		this.size++
	}
}


func (this *MinStack) Pop()  {
	if this.size!=0{
		this.s = append(this.s[:this.size-1])
		if this.size-1==this.min{
			minum := 1<<32-1
			this.min=0
			var i int
			for i = 0;i<this.size-1;i++{
				if(this.s[i]<minum){
					minum =this.s[i]
					this.min = i
				}
			}
		}
		this.size--
	}
}


func (this *MinStack) Top() int {
	if this.size!=0{
		return int(this.s[this.size-1])
	}else {
		return -1
	}
}


func (this *MinStack) GetMin() int {
	if this.size!=0{
		return int(this.s[this.min])
	}else {
		return -1
	}
}

func main()  {
	var m MinStack = Constructor()
	m.Push(2147483646)
	m.Push(2147483646)
	m.Push(2147483647)
	m.Top()
	m.Pop()
	m.GetMin()
	m.Pop()
	m.GetMin()
	m.Pop()
	m.Push(2147483647)
	fmt.Println(m.Top())
	m.GetMin()
	m.Push(-2147483648)
	m.Top()
	m.GetMin()
	m.Pop()
	m.GetMin()
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
