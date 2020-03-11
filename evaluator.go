package calculator

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
			}
		}
	}
	ne, err := s.pop()
	if err != nil || !s.empty() {
		return -1, err
	}
	n, _ := ne.(int)
	return n, nil
}
