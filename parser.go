package calculator

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type parser struct {
	rs  []rune
	pos int
}

type expr interface{}

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

func isLit(r rune) bool {
	switch r {
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		fallthrough
	case '0':
		return true
	default:
		return false
	}
}

func isExpr(r rune) bool {
	return isBlank(r) || isOp(r) || isLit(r)
}

type eol bool

func (p *parser) next() (expr, eol, error) {
	if p.isEOF() {
		return nil, eol(true), nil
	}
	r := p.rs[p.pos]
	if !isExpr(r) {
		ps := p.pos
		p.pos++
		return nil, eol(false), fmt.Errorf("invalid character %c at column %d", r, ps)
	}
	for ; isBlank(r); r = p.rs[p.pos] {
		if p.hasNext() {
			p.pos++
		} else {
			return nil, eol(true), nil
		}
	}
	if isOp(r) {
		p.pos++
		return op(r), eol(false), nil
	}
	s := ""
	for r = p.rs[p.pos]; isLit(r); r = p.rs[p.pos] {
		s += string(r)
		if p.hasNext() {
			p.pos++
		} else {
			break
		}
	}
	k, err := strconv.Atoi(s)
	if err != nil {
		return nil, eol(false), err
	}
	return lit(k), eol(false), nil
}