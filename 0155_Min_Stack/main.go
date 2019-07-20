package Min_Stack

/**
ref:https://leetcode.com/problems/min-stack/

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
Example:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
*/
type MinStack struct {
	stackArray []int
	miniArray  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.stackArray = append(this.stackArray, x)

	if len(this.miniArray) == 0 || x <= this.miniArray[len(this.miniArray)-1] {
		this.miniArray = append(this.miniArray, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.stackArray) > 0 {
		pop := this.stackArray[len(this.stackArray)-1]
		this.stackArray = this.stackArray[0 : len(this.stackArray)-1]
		if len(this.miniArray) > 0 &&
			pop == this.miniArray[len(this.miniArray)-1] {
			this.miniArray = this.miniArray[0 : len(this.miniArray)-1]
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.stackArray) > 0 {
		return this.stackArray[len(this.stackArray)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.miniArray) > 0 {
		return this.miniArray[len(this.miniArray)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
