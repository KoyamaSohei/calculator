package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	assert.Equal(t, 1, eval("1"))
}
