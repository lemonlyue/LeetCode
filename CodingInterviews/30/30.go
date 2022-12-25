package _30

// MinStack
// 包含min函数的栈
// https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/description/
type MinStack struct {
	Stack    []int
	StackMin []int
}

// Constructor
/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)
	if len(this.StackMin) == 0 {
		this.StackMin = append(this.StackMin, x)
	} else {
		if this.StackMin[len(this.StackMin)-1] > x {
			this.StackMin = append(this.StackMin, x)
		} else {
			this.StackMin = append(this.StackMin, this.StackMin[len(this.StackMin)-1])
		}
	}
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.StackMin = this.StackMin[:len(this.StackMin)-1]
}

func (this *MinStack) Top() int {
	value := this.Stack[len(this.Stack)-1]
	return value
}

func (this *MinStack) Min() int {
	return this.StackMin[len(this.StackMin)-1]
}
