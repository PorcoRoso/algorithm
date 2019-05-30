package stack

type MyQueue struct {
	inStack []int
	outStack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	inStack := make([]int, 0)
	outStack := make([]int, 0)

	return MyQueue{
		inStack: inStack,
		outStack: outStack,
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.inStack = append(this.inStack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		for i := len(this.inStack); i > 0; i-- {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	res := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) != 0 {
		return this.outStack[len(this.outStack)-1]
	} else {
		return this.inStack[0]
	}
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.inStack) > 0 || len(this.outStack) > 0 {
		return false
	} else {
		return true
	}
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
