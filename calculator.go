package calculator

func eval(s string) (int, error) {
	e, err := parse(s)
	if err != nil {
		return -1, err
	}
	e, err = shuntingyard(e)
	if err != nil {
		return -1, err
	}
	return evalPostfix(e)
}
