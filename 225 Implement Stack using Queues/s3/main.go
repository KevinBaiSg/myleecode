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
	queue *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{queue: &Queue{array: make([]int, 0), size: 0}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue.Push(x)

	size := this.queue.size
	for size > 1 {
		x, _ := this.queue.Pop()
		this.queue.Push(x)
		size--
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.queue.size > 0 {
		x, _ := this.queue.Pop()
		return x
	} else {
		return -1
	}
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.queue.size > 0 {
		x, _ := this.queue.front()
		return x
	} else {
		return -1
	}
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue.size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
