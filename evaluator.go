package calculator

import "fmt"

var precedence = map[op]int{
	op('+'): 4,
	op('-'): 4,
	op('*'): 5,
	op('/'): 5}

func isHighPrecedence(o1 op, o2 op) bool {
	return precedence[o1] > precedence[o2]
}

func shuntingyard(ex []expr) ([]expr, error) {
	out := make([]expr, 0)
	s := newStack()
	for _, e := range ex {
		switch c := e.(type) {
		case lit:
			out = append(out, c)
		case bra:
			switch c {
			case bra('('):
				s.push(c)
			case bra(')'):
				for r, err := s.pop(); r != bra('('); r, err = s.pop() {
					if err != nil {
						return nil, fmt.Errorf("mismatched parentheses")
					}
					out = append(out, r)
				}
			}
		case op:
			for !s.empty() {
				r, err := s.top()
				if err != nil {
					return nil, err
				}
				_, ok := r.(bra)
				if ok {
					break
				}
				if !isHighPrecedence(c, r.(op)) {
					r, err = s.pop()
					if err != nil {
						return nil, err
					}
					out = append(out, r)
				} else {
					break
				}
			}
			s.push(c)
		}
	}
	for !s.empty() {
		r, err := s.pop()
		_, ok := r.(bra)
		if ok {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		if err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, nil
}

func evalPostfix(ex []expr) (int, error) {
	s := newStack()
	for _, e := range ex {
		switch c := e.(type) {
		case lit:
			s.push(c)
		case op:
			be, err := s.pop()
			if err != nil {
				return -1, fmt.Errorf("invalid operation at %c", c)
			}
			b, _ := be.(lit)
			ae, err := s.pop()
			if err != nil {
				return -1, fmt.Errorf("invalid operation at %c", c)
			}
			a, _ := ae.(lit)
			switch c {
			case op('+'):
				s.push(a + b)
			case op('-'):
				s.push(a - b)
			case op('*'):
				s.push(a * b)
			case op('/'):
				s.push(a / b)
			}
		}
	}
	ne, err := s.pop()
	if err != nil || !s.empty() {
		return -1, fmt.Errorf("invalid operation")
	}
	nl, _ := ne.(lit)
	return int(nl), nil
}
