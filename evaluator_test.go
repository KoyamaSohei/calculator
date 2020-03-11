package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvalPostfix(t *testing.T) {
	k, err := evalPostfix([]expr{lit(1), lit(1), op('+')})
	assert.Nil(t, err)
	assert.Equal(t, 2, k)
	k, err = evalPostfix([]expr{lit(1), lit(3), lit(1), op('+'), op('+')})
	assert.Nil(t, err)
	assert.Equal(t, 5, k)
	_, err = evalPostfix([]expr{lit(1), lit(3), lit(1), op('+'), op('+'), op('+')})
	assert.NotNil(t, err)
	_, err = evalPostfix([]expr{op('+'), op('+'), lit(0), op('+')})
	assert.NotNil(t, err)
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
