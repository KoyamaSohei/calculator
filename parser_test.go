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
	assert.Equal(t, err, fmt.Errorf("invalid character & at column 1"))
	assert.Equal(t, n, nil)
	assert.Equal(t, e, eol(false))
	n, e, err = p.next()
	assert.Equal(t, err, nil)
	assert.Equal(t, n, nil)
	assert.Equal(t, e, eol(true))
}

func TestParse(t *testing.T) {
	cases := []testpairAdv{
		testpairAdv{"2", []expr{lit(2)}},
		testpairAdv{"3+45+67+890", []expr{lit(3), op('+'), lit(45), op('+'), lit(67), op('+'), lit(890)}},
		testpairAdv{" 1+ 22+333 +4444 ", []expr{lit(1), op('+'), lit(22), op('+'), lit(333), op('+'), lit(4444)}},
		testpairAdv{"++++1+++++", []expr{op('+'), op('+'), op('+'), op('+'), lit(1), op('+'), op('+'), op('+'), op('+'), op('+')}}}
	for _, c := range cases {
		t.Log(c.s)
		i, err := parse(c.s)
		assert.Equal(t, nil, err)
		assert.Equal(t, len(c.i), len(i))
		assert.Equal(t, c.i, i)
	}
}
