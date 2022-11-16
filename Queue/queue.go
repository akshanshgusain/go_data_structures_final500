package Queue

import (
	"container/list"
	"fmt"
)

type Queue struct {
	Queue *list.List
}

func (s *Queue) Size() int {
	return s.Queue.Len()
}

func (s *Queue) Empty() bool {
	return s.Queue.Len() == 0
}

//Push element at the front of the list: O(1)
func (s *Queue) Push(value string) {
	s.Queue.PushBack(value)
}

//Pop element from the front if the list: O(1)
func (s *Queue) Pop() error {
	if s.Queue.Len() > 0 {
		e := s.Queue.Front()
		s.Queue.Remove(e)
		return nil
	}
	return fmt.Errorf("queue is empty")
}

// Front Get the top element: O(1)
func (s *Queue) Front() (string, error) {
	if s.Queue.Len() > 0 {
		if val, ok := s.Queue.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("queue datatype is incorrect")
	}
	return "", fmt.Errorf("queue is empty")
}
