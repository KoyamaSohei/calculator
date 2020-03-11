package calculator

import (
	"errors"
	"sync"
)

type stack struct {
	lock sync.Mutex
	s    []lit
}

func newStack() *stack {
	return &stack{sync.Mutex{}, make([]lit, 0)}
}

func (s *stack) push(v lit) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) pop() (lit, error) {
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

func (s *stack) empty() bool {
	l := len(s.s)
	if l == 0 {
		return true
	}
	return false
}
