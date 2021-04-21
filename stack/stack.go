package stack

import "errors"

type Stack []interface{}

func New(len int) Stack {
	return make(Stack, 0, len)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack) IsFull() bool {
	return len(s) == cap(s)
}

func (s Stack) Cap() int {
	return cap(s)
}

func (s *Stack) Push(element interface{}) (bool, error) {
	if s.IsFull() {
		return false, errors.New("the stack is full!")
	}
	*s = append(*s, element)
	return true, nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("the stack is empty!")
	}
	thisStack := *s
	element := thisStack[len(thisStack)-1]
	*s = thisStack[:len(thisStack)-1]
	return element, nil
}

func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("the stack is empty!")
	}
	thisStack := *s
	return thisStack[len(thisStack)-1], nil
}
