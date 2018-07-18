package brackets

type Stack struct {
	index int
	stack []uint8
}

func CreateStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(val uint8) {
	stack.index++
	stack.stack = append(stack.stack, val)
}

func (stack *Stack) Pop() (res uint8) {
	stack.index--
	res = stack.stack[stack.index]
	stack.stack = stack.stack[:stack.index]
	return
}

func is_brackets(char uint8, brackets string) (res bool, index int) {
	res = false
	for index = range brackets {
		if char == brackets[index] {
			res = true
			return
		}
	}
	return
}

func Bracket(brackets string) (bool, error) {

	stack := CreateStack()
	open_br := "{[("
	close_br := "}])"
	var res bool
	var index int

	for i, val := range brackets {
		if res, index = is_brackets(uint8(val), open_br); res {
			stack.Push(uint8(val))
		} else if res, index = is_brackets(uint8(val), close_br); res {
			if i > 0 {
				if stack.stack[stack.index-1] == open_br[index] {
					stack.Pop()
				} else {
					return false, nil
				}
			} else {
				return false, nil
			}
		} else {
			return false, nil
		}
	}
	if stack.index > 0 {
		return false, nil
	}
	return true, nil
}
