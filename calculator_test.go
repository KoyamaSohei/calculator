package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testIO struct {
	i string
	o int
}

func TestCalculator(t *testing.T) {
	ios := []testIO{
		testIO{"1", 1},
		testIO{"1+1", 2},
		testIO{"2*100", 200},
		testIO{"2+3*6", 20},
		testIO{" 2+100-50 *4 / 5", 62},
		testIO{"2*(1+5*10)", 102},
		testIO{"5+((((((2))))))*7", 19},
		testIO{"3+(1+4*4)*5", 88}}
	for _, io := range ios {
		o, err := Eval(io.i)
		assert.Nil(t, err)
		assert.Equal(t, io.o, o)
	}
}

type testIOErr struct {
	i string
	o error
}

func TestCalculatorError(t *testing.T) {
	ios := []testIOErr{
		testIOErr{"1+", fmt.Errorf("invalid operation at +")},
		testIOErr{"1+(2*3", fmt.Errorf("mismatched parentheses")},
		testIOErr{"100/0", fmt.Errorf("division by zero")},
		testIOErr{"+2*3", fmt.Errorf("invalid operation at +")}}
	for _, io := range ios {
		o, err := Eval(io.i)
		assert.Equal(t, -1, o)
		assert.Equal(t, io.o, err)
	}
}
