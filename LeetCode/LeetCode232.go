package main

import (
	"container/list"
)

type MyQueue struct {
	stack Stack
}

type Stack struct {
	length int
	list list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	q:=MyQueue{}
	q.stack.list = list.List{}
	q.stack.length=0
	return q
}

func (this *Stack) Push(x int)  {
	this.list.PushBack(x)
	this.length++
}

func (this *Stack) Pop() int {
	r := this.list.Back()
	this.list.Remove(r)
	this.length--
	return r.Value.(int)
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	st:=Stack{}
	for true{
		if this.stack.length==0 {break}
		st.Push(this.stack.Pop())
	}
	st.Push(x)
	for true{
		if st.length==0 {break}
		this.stack.Push(st.Pop())
	}
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	return this.stack.Pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	res:= this.stack.Pop()
	this.stack.Push(res)
	return res
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack.length==0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */