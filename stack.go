package wikiwiki

type stack[T comparable] struct {
	items []T
}

func (receiver *stack[T]) Peek() (T, bool) {
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

func (receiver *stack[T]) Pop() (T, bool) {
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

func (receiver *stack[T]) Push(value T) {
	if nil == receiver {
		return
	}

	receiver.items = append(receiver.items, value)
}

func (receiver *stack[T]) TopEqual(value T) bool {
	if nil == receiver {
		return false
	}

	top, found := receiver.Peek()
	if !found {
		return false
	}

	return top == value
}
