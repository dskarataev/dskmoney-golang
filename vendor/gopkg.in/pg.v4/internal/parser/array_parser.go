package parser

import (
	"bytes"
	"fmt"
)

type ArrayParser struct {
	*Parser

	stickyErr error
}

func NewArrayParser(b []byte) *ArrayParser {
	var err error
	if len(b) < 2 || b[0] != '{' || b[len(b)-1] != '}' {
		err = fmt.Errorf("pg: can't parse array: %s", string(b))
	} else {
		b = b[1 : len(b)-1]
	}
	return &ArrayParser{
		Parser: New(b),

		stickyErr: err,
	}
}

func (p *ArrayParser) NextElem() ([]byte, error) {
	if p.stickyErr != nil {
		return nil, p.stickyErr
	}

	switch c := p.Peek(); c {
	case '"':
		p.Advance()
		b := p.readSubstring()
		p.Skip(',')
		return b, nil
	case '{':
		b := p.readElem()
		if b != nil {
			b = append(b, '}')
		}
		p.Skip(',')
		return b, nil
	default:
		b, _ := p.ReadSep(',')
		if bytes.Equal(b, pgNull) {
			b = nil
		}
		return b, nil
	}

	return nil, fmt.Errorf("pg: can't parse array: %q", p.Bytes())
}

func (p *ArrayParser) readElem() []byte {
	var b []byte
	for p.Valid() {
		c := p.Read()
		switch c {
		case '"':
			b = append(b, '"')
			for {
				bb, ok := p.ReadSep('"')
				b = append(b, bb...)
				stop := len(b) > 0 && b[len(b)-1] != '\\'
				if ok {
					b = append(b, '"')
				}
				if stop {
					break
				}
			}
		case '}':
			return b
		default:
			b = append(b, c)
		}
	}
	return b
}
