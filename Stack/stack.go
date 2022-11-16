package Stack

import (
	"container/list"
	"fmt"
)

type Stack struct {
	stack *list.List
}

//Push element at the front of the list: O(1)
func (s *Stack) Push(value string) {
	s.stack.PushFront(value)
}

//Pop element from the front if the list: O(1)
func (s *Stack) Pop() error {
	if s.stack.Len() > 0 {
		e := s.stack.Front()
		s.stack.Remove(e)
		return nil
	}
	return fmt.Errorf("stack is empty")
}

// Top Get the top element: O(1)
func (s *Stack) Top() (string, error) {
	if s.stack.Len() > 0 {
		if val, ok := s.stack.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("stack datatype is incorrect")
	}
	return "", fmt.Errorf("stack is empty")
}

func (s *Stack) Size() int {
	return s.stack.Len()
}

func (s *Stack) Empty() bool {
	return s.stack.Len() == 0
}

// Second Implementation

//type Stack []string
//
//// IsEmpty :IsEmpty check if stack is empty
//func (s *Stack) IsEmpty() bool {
//	return len(*s) == 0
//}
//
//// Push a new value onto the stack
//func (s *Stack) Push(str string) {
//	*s = append(*s, str) // Simply append the new value to the end of the stack
//}
//
//// Pop Remove and return top element of stack. Return false if stack is empty.
//func (s *Stack) Pop() (string, bool) {
//	if s.IsEmpty() {
//		return "", false
//	} else {
//		index := len(*s) - 1   // Get the index of the top most element.
//		element := (*s)[index] // Index into the slice and obtain the element.
//		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
//		return element, true
//	}
//}
