package main

import "fmt"

type MyStack struct {
	queue Queue
}

type Queue struct {
	content []int
	length int
}
func QueueConstructor() Queue {
	var q Queue
	q.content = make([]int,0)
	q.length = 0
	return q
}

func (this *Queue) Push(x int)  {
	this.content = append(this.content, x)
	this.length++
}

func (this *Queue) Pop() int {
	res := this.content[len(this.content)-1]
	this.content = this.content[0:len(this.content)-1]
	this.length--
	return res
}

func (this *Queue) Top() int {
	return this.content[this.length-1]
}

func (this *Queue) Empty() bool {
	  return this.length==0
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	var stack MyStack
	stack.queue = QueueConstructor()
	return stack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	if this.queue.length==0 {this.queue.Push(x);return }
	this.queue.Push(x)
	for i:=0;i<this.queue.length-1;i++ {
		this.queue.Push(this.queue.Pop())
	}
	return
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.queue.Pop()
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue.Top()
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue.length==0
}

func main()  {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
}