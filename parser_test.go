package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOp(t *testing.T) {
	assert.Equal(t, true, isOp('+'))
	assert.Equal(t, false, isOp('0'))
}

type testpair struct {
	s string
	i expr
}

func TestNext(t *testing.T) {
	cases := []testpair{
		testpair{"1", lit(1)},
		testpair{"11", lit(11)},
		testpair{" 1", lit(1)},
		testpair{"1 ", lit(1)},
		testpair{" 11 ", lit(11)},
		testpair{"+", op('+')},
		testpair{" +", op('+')},
		testpair{"+ ", op('+')},
		testpair{" + ", op('+')}}
	for _, c := range cases {
		p := newParser(c.s)
		n, _, err := p.next()
		if err != nil {
			t.Log("\"" + c.s + "\"")
			t.Fatal(err)
		}
		assert.Equal(t, c.i, n)
	}
}

type testpairAdv struct {
	s string
	i []expr
}

func TestNextAdv(t *testing.T) {
	cases := []testpairAdv{
		testpairAdv{"1+1", []expr{lit(1), op('+'), lit(1)}},
		testpairAdv{"1 + 1", []expr{lit(1), op('+'), lit(1)}},
		testpairAdv{" 1+    1", []expr{lit(1), op('+'), lit(1)}},
		testpairAdv{" 1     +1     ", []expr{lit(1), op('+'), lit(1)}}}
	for _, c := range cases {
		p := newParser(c.s)
		for _, d := range c.i {
			n, _, err := p.next()
			if err != nil {
				t.Log("\"" + c.s + "\"")
				t.Log(d)
				t.Fatal(err)
			}
			assert.Equal(t, d, n)
		}
	}
}

func TestNextError(t *testing.T) {
	p := newParser("1&")
	n, e, err := p.next()
	assert.Equal(t, err, nil)
	assert.Equal(t, n, lit(1))
	assert.Equal(t, e, eol(false))
	n, e, err = p.next()
	assert.Equal(t, err, fmt.Errorf("invalid character & at column1"))
	assert.Equal(t, n, nil)
	assert.Equal(t, e, eol(false))
	n, e, err = p.next()
	assert.Equal(t, err, nil)
	assert.Equal(t, n, nil)
	assert.Equal(t, e, eol(true))
}
