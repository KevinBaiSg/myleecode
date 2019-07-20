package Implement_Queue_using_Stacks

type MyStack struct {
	array 	[]int
	size 	int
}

func (this *MyStack) Push(x int)  {
	this.array = append(this.array, x)
	this.size ++
}

func (this *MyStack) Pop() int {
	if this.size > 0 {
		x := this.array[this.size - 1]
		this.array = this.array[:this.size - 1]
		this.size --
		return x
	}
	return -1
}

func (this *MyStack) Peek() int {
	if this.size > 0 {
		return this.array[this.size - 1]
	}
	return -1
}

func (this *MyStack) Empty() bool {
	return this.size == 0
}


type MyQueue struct {
	stackPush *MyStack
	stackPop *MyStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stackPush:&MyStack{array:make([]int, 0), size:0},
		stackPop:&MyStack{array:make([]int, 0), size:0},
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	if !this.stackPop.Empty() {
		size := this.stackPop.size
		for size > 0 {
			this.stackPush.Push(this.stackPop.Pop())
			size --
		}
	}
	this.stackPush.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if !this.stackPop.Empty() {
		return this.stackPop.Pop()
	}

	if this.stackPush.Empty() {
		return -1
	}

	size := this.stackPush.size
	for size > 1 {
		this.stackPop.Push(this.stackPush.Pop())
		size --
	}
	return this.stackPush.Pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if !this.stackPop.Empty() {
		return this.stackPop.Peek()
	}

	if this.stackPush.Empty() {
		return -1
	}

	size := this.stackPush.size
	for size > 0 {
		this.stackPop.Push(this.stackPush.Pop())
		size --
	}
	return this.stackPop.Peek()
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return (this.stackPush.size == 0) && (this.stackPop.size == 0)
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */