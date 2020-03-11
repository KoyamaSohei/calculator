package calculator

func evalPostfix(ex []expr) (int, error) {
	s := newStack()
	for _, e := range ex {
		switch c := e.(type) {
		case lit:
			s.push(c)
		case op:
			b, err := s.pop()
			if err != nil {
				return -1, err
			}
			a, err := s.pop()
			if err != nil {
				return -1, err
			}
			switch c {
			case op('+'):
				s.push(a + b)
			}
		}
	}
	n, err := s.pop()
	if err != nil || !s.empty() {
		return -1, err
	}
	return int(n), nil
}
