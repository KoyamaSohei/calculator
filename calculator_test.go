package calculator

import (
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
		testIO{" 2+100-50 *4 / 5", 62}}
	for _, io := range ios {
		o, err := Eval(io.i)
		assert.Nil(t, err)
		assert.Equal(t, io.o, o)
	}
}
