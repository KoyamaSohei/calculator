package calculator

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
		case op:
			for !s.empty() {
				r, err := s.top()
				if err != nil {
					return nil, err
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
				return -1, err
			}
			b, _ := be.(lit)
			ae, err := s.pop()
			if err != nil {
				return -1, err
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
		return -1, err
	}
	nl, _ := ne.(lit)
	return int(nl), nil
}
