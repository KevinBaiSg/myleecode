package s1

import (
	"errors"
)

type Queue struct {
	array []int
	size  int
}

func (this *Queue) Push(x int) {
	this.array = append(this.array, x)
	this.size++
}

func (this *Queue) Pop() (int, error) {
	if this.size > 0 {
		x := this.array[0]
		this.array = this.array[1:this.size]
		this.size--
		return x, nil
	}
	return 0, errors.New("queue is empty")
}

func (this *Queue) back() (int, error) {
	if this.size > 0 {
		return this.array[this.size-1], nil
	}
	return 0, errors.New("queue is empty")
}
func (this *Queue) front() (int, error) {
	if this.size > 0 {
		return this.array[0], nil
	}
	return 0, errors.New("queue is empty")
}

func (this *Queue) Empty() bool {
	return this.size == 0
}

type MyStack struct {
	queueL *Queue
	queueR *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{queueL: &Queue{array: make([]int, 0), size: 0}, queueR: &Queue{array: make([]int, 0), size: 0}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.queueL.Empty() {
		this.queueR.Push(x)
	} else {
		this.queueL.Push(x)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	var queuePush, queuePop *Queue
	if !this.queueL.Empty() {
		queuePop = this.queueL
		queuePush = this.queueR
	} else if !this.queueR.Empty() {
		queuePop = this.queueR
		queuePush = this.queueL
	} else {
		return -1
	}

	size := queuePop.size
	for size != 1 {
		x, _ := queuePop.Pop()
		queuePush.Push(x)
		size--
	}
	x, _ := queuePop.Pop()
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if !this.queueL.Empty() {
		x, _ := this.queueL.back()
		return x
	} else if !this.queueR.Empty() {
		x, _ := this.queueR.back()
		return x
	} else {
		return -1
	}
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return (this.queueL.Empty()) && (this.queueR.Empty())
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
