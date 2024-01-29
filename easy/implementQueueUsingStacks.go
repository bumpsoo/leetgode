package easy

// https://leetcode.com/problems/implement-queue-using-stacks
type IntStack []int

func (i *IntStack) Push(v int) {
	(*i) = append((*i), v)
}

func (i *IntStack) Pop() int {
	l := len(*i) - 1
	ret := (*i)[l]
	(*i) = (*i)[:l]
	return ret
}

func (i *IntStack) Peek() int {
	return (*i)[len(*i)-1]
}

func (i *IntStack) Len() int {
	return len(*i)
}

type MyQueue struct {
	A, B IntStack
}

func Constructor() MyQueue {
	return MyQueue{A: IntStack{}}
}

func (this *MyQueue) Push(x int) {
	if this.A.Len() >= 1 {
		this.B.Push(x)
	} else {
		this.A.Push(x)
	}
}

func (this *MyQueue) fix() {
	switch this.B.Len() {
	case 0:
	case 1:
		this.A.Push(this.B.Pop())
	default:
		v := this.B.Pop()
		this.fix()
		this.B.Push(v)
	}
}

func (this *MyQueue) Pop() int {
	ret := this.A.Pop()
	this.fix()
	return ret
}

func (this *MyQueue) Peek() int {
	return this.A.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.A.Len() == 0 && this.B.Len() == 0
}
