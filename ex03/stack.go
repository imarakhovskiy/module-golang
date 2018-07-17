package stack

type Stack struct {
	index int
	stack []int
}

func New() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(val int) {
	stack.index++
	stack.stack = append(stack.stack, val)
}

func (stack *Stack) Pop() (res int) {
	stack.index--
	res = stack.stack[stack.index]
	stack.stack = stack.stack[:stack.index]
	return
}
