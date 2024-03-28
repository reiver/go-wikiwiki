package stack

// Stack implements a stack.
type Stack[T comparable] struct {
	items []T
}

func (receiver *Stack[T]) Peek() (T, bool) {
	var nothing T

	if nil == receiver {
		return nothing, false
	}

	var length int = len(receiver.items)
	var last int = length -1

	if last < 0 {
		return nothing, false
	}

	var value T

	value = receiver.items[last]

	return value, true
}

func (receiver *Stack[T]) Pop() (T, bool) {
	var nothing T

	if nil == receiver {
		return nothing, false
	}

	var length int = len(receiver.items)
	var last int = length -1

	if last < 0 {
		return nothing, false
	}

	var value T

	value, receiver.items = receiver.items[last], receiver.items[:last]

	return value, true
}

func (receiver *Stack[T]) Push(value T) {
	if nil == receiver {
		return
	}

	receiver.items = append(receiver.items, value)
}

func (receiver *Stack[T]) TopEqual(value T) bool {
	if nil == receiver {
		return false
	}

	top, found := receiver.Peek()
	if !found {
		return false
	}

	return top == value
}
