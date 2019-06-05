package leetcode

import (
	"bytes"
)

func removeDuplicates(S string) string {
	if len(S) <= 1 {
		return S
	}
	stack := NewStack()
	for _, v := range S {
		if stack.len() == 0 {
			stack.push(byte(v))
		} else {
			if stack.top() == byte(v) {
				stack.pop()
			} else {
				stack.push(byte(v))
			}
		}
	}
	return stack.String()
}

type Stack struct {
	value []byte
}

func NewStack() *Stack {
	return &Stack{
		value: make([]byte, 0),
	}
}

func (s *Stack) push(b byte) {
	s.value = append(s.value, b)
}

func (s *Stack) pop() byte {
	tmp := s.value[len(s.value)-1]
	s.value = s.value[:len(s.value)-1]
	return tmp
}

func (s *Stack) top() byte {
	return s.value[len(s.value)-1]
}

func (s *Stack) len() int {
	return len(s.value)
}

func (s *Stack) String() string {
	var str bytes.Buffer
	str.Write(s.value)
	return str.String()
}
