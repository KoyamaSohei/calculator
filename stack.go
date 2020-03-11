package calculator

import (
	"errors"
	"sync"
)

type stack struct {
	lock sync.Mutex
	s    []expr
}

func newStack() *stack {
	return &stack{sync.Mutex{}, make([]expr, 0)}
}

func (s *stack) push(v expr) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) pop() (expr, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *stack) top() (expr, error) {
	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	return res, nil
}

func (s *stack) empty() bool {
	l := len(s.s)
	if l == 0 {
		return true
	}
	return false
}
