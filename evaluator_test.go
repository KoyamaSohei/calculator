package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuntingyard(t *testing.T) {
	o, err := shuntingyard([]expr{lit(1), op('+'), lit(1)})
	assert.Nil(t, err)
	assert.Equal(t, []expr{lit(1), lit(1), op('+')}, o)
	o, err = shuntingyard([]expr{lit(10), op('+'), lit(2), op('*'), lit(3)})
	assert.Nil(t, err)
	assert.Equal(t, []expr{lit(10), lit(2), lit(3), op('*'), op('+')}, o)
	o, err = shuntingyard([]expr{lit(4), op('/'), lit(2), op('-'), lit(10)})
	assert.Nil(t, err)
	assert.Equal(t, []expr{lit(4), lit(2), op('/'), lit(10), op('-')}, o)
	o, err = shuntingyard([]expr{lit(1), op('+'), lit(2), op('-'), lit(3), op('+'), lit(4)})
	assert.Nil(t, err)
	assert.Equal(t, []expr{lit(1), lit(2), op('+'), lit(3), op('-'), lit(4), op('+')}, o)
	o, err = shuntingyard([]expr{lit(1), op('+')})
	assert.Nil(t, err)
	assert.Equal(t, []expr{lit(1), op('+')}, o)
}

func TestShuntingyardAdv(t *testing.T) {
	o, err := shuntingyard([]expr{lit(9), op('*'), bra('('), lit(1), op('+'), lit(2), op('*'), lit(10), bra(')')})
	assert.Nil(t, err)
	assert.Equal(t, []expr{lit(9), lit(1), lit(2), lit(10), op('*'), op('+'), op('*')}, o)
	o, err = shuntingyard([]expr{lit(3), op('+'), bra('('), lit(1), op('+'), lit(4), op('*'), lit(4), bra(')'), op('*'), lit(5)})
	assert.Nil(t, err)
	assert.Equal(t, []expr{lit(3), lit(1), lit(4), lit(4), op('*'), op('+'), lit(5), op('*'), op('+')}, o)
	o, err = shuntingyard([]expr{bra(')')})
	assert.Equal(t, fmt.Errorf("mismatched parentheses"), err)
	assert.Nil(t, o)
	o, err = shuntingyard([]expr{bra('(')})
	assert.Equal(t, fmt.Errorf("mismatched parentheses"), err)
	assert.Nil(t, o)
	o, err = shuntingyard([]expr{lit(1), bra('('), op('+'), lit(2)})
	assert.Equal(t, fmt.Errorf("mismatched parentheses"), err)
	assert.Nil(t, o)
	o, err = shuntingyard([]expr{op('*'), bra('('), op('+'), lit(2)})
	assert.Equal(t, fmt.Errorf("mismatched parentheses"), err)
	assert.Nil(t, o)
	o, err = shuntingyard([]expr{lit(1), bra(')'), op('*'), lit(2)})
	assert.Equal(t, fmt.Errorf("mismatched parentheses"), err)
	assert.Nil(t, o)
	o, err = shuntingyard([]expr{lit(1), bra(')'), lit(2), lit(2)})
	assert.Equal(t, fmt.Errorf("mismatched parentheses"), err)
	assert.Nil(t, o)
}

func TestEvalPostfix(t *testing.T) {
	k, err := evalPostfix([]expr{lit(1), lit(1), op('+')})
	assert.Nil(t, err)
	assert.Equal(t, 2, k)
	k, err = evalPostfix([]expr{lit(1), lit(3), lit(1), op('+'), op('+')})
	assert.Nil(t, err)
	assert.Equal(t, 5, k)
	_, err = evalPostfix([]expr{lit(1), lit(3), lit(1), op('+'), op('+'), op('+')})
	assert.Equal(t, fmt.Errorf("invalid operation at +"), err)
	_, err = evalPostfix([]expr{op('+'), op('+'), lit(0), op('+')})
	assert.Equal(t, fmt.Errorf("invalid operation at +"), err)
	k, err = evalPostfix([]expr{lit(1), lit(3), op('-')})
	assert.Nil(t, err)
	assert.Equal(t, -2, k)
	k, err = evalPostfix([]expr{lit(2), lit(9), op('*')})
	assert.Nil(t, err)
	assert.Equal(t, 18, k)
	k, err = evalPostfix([]expr{lit(8), lit(4), op('/')})
	assert.Nil(t, err)
	assert.Equal(t, 2, k)
	k, err = evalPostfix([]expr{lit(9), lit(4), op('/')})
	assert.Nil(t, err)
	assert.Equal(t, 2, k)
}

func TestEvalPostfixAdv(t *testing.T) {
	k, err := evalPostfix([]expr{lit(1), lit(3), lit(10), op('*'), op('+')})
	assert.Nil(t, err)
	assert.Equal(t, 31, k)
	k, err = evalPostfix([]expr{lit(3), lit(7), lit(5), op('-'), op('*')})
	assert.Nil(t, err)
	assert.Equal(t, 6, k)
	k, err = evalPostfix([]expr{lit(4), lit(10), op('-')})
	assert.Nil(t, err)
	assert.Equal(t, -6, k)
	k, err = evalPostfix([]expr{lit(4), lit(5), op('-'), lit(100), op('/')})
	assert.Nil(t, err)
	assert.Equal(t, 0, k)
}
