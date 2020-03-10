package calculator

import (
	"strconv"
	"unicode/utf8"
)

type parser struct {
	rs  []rune
	pos int
}

type node interface{}

type lit int

type op rune

func newParser(s string) parser {
	rs := make([]rune, 0)
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		rs = append(rs, r)
		s = s[size:]
	}
	return parser{rs, 0}
}

func (p parser) isEOF() bool {
	return p.pos >= len(p.rs)
}

func (p parser) hasNext() bool {
	return p.pos+1 < len(p.rs)
}

func isBlank(r rune) bool {
	return r == ' '
}

func isOp(r rune) bool {
	return r == '+'
}

type eof bool

func (p *parser) next() (node, eof, error) {
	if p.isEOF() {
		return nil, eof(true), nil
	}
	var r rune
	for r = p.rs[p.pos]; isBlank(r); r = p.rs[p.pos] {
		if p.hasNext() {
			p.pos++
		} else {
			return nil, eof(true), nil
		}
	}
	if isOp(r) {
		p.pos++
		return op(r), eof(false), nil
	}
	s := ""
	for r = p.rs[p.pos]; !isBlank(r) && !isOp(r); r = p.rs[p.pos] {
		s += string(r)
		if p.hasNext() {
			p.pos++
		} else {
			break
		}
	}
	k, err := strconv.Atoi(s)
	if err != nil {
		return nil, eof(false), err
	}
	return lit(k), eof(false), nil
}
