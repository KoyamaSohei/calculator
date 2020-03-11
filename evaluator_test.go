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
}
