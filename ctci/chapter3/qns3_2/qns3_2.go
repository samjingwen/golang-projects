package qns3_2

type Stack struct {
	values    []int
	minValues []int
}

func (stack *Stack) append(value int) {
	if len(stack.values) != 0 {
		if minVal := stack.minValues[len(stack.minValues)-1]; minVal > value {
			stack.minValues = append(stack.minValues, value)
		} else {
			stack.minValues = append(stack.minValues, minVal)
		}
	} else {
		stack.minValues = append(stack.minValues, value)
	}
	stack.values = append(stack.values, value)
}

func (stack *Stack) pop() int {
	val := stack.values[len(stack.values)-1]
	stack.values = stack.values[:len(stack.values)-1]
	stack.minValues = stack.minValues[:len(stack.minValues)-1]
	return val
}

func (stack *Stack) min() int {
	return stack.minValues[len(stack.minValues)-1]
}
