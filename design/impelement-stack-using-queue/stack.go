package stack

import "github.com/lephuocmy668/challengers/queue"

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement Stack using Queues.
// https://leetcode.com/problems/implement-stack-using-queues/submissions/
type MyStack struct {
	q1 *queue.Queue
	q2 *queue.Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		q1: queue.NewQueue(),
		q2: queue.NewQueue(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q2.Enqueue(x)
	for !this.q1.IsEmpty() {
		val := this.q1.Dequeue()
		this.q2.Enqueue(*val)
	}

	this.q1, this.q2 = this.q2, this.q1
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.q1.IsEmpty() {
		return -1
	}
	return *this.q1.Dequeue()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.q1.IsEmpty() {
		return -1
	}
	return this.q1.Peek()
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q1.IsEmpty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
