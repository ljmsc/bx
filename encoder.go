package bx

import (
	"bytes"
	"fmt"
)

type E struct {
	values []ValueType
}

func Encoder() *E {
	e := E{
		values: []ValueType{},
	}
	return &e
}

func (b *E) Write(vt ValueType) *E {
	if b.values == nil {
		b.values = []ValueType{}
	}
	b.values = append(b.values, vt)
	return b
}

func (b *E) Int8(val int8) *E {
	return b.Write(&Number{value: val})
}

func (b *E) Int16(val int16) *E {
	return b.Write(&Number{value: val})
}

func (b *E) Int32(val int32) *E {
	return b.Write(&Number{value: val})
}

func (b *E) Int64(val int64) *E {
	return b.Write(&Number{value: val})
}

func (b *E) Uint8(val uint8) *E {
	return b.Write(&Number{value: val})
}

func (b *E) Uint16(val uint16) *E {
	return b.Write(&Number{value: val})
}

func (b *E) Uint32(val uint32) *E {
	return b.Write(&Number{value: val})
}

func (b *E) Uint64(val uint64) *E {
	return b.Write(&Number{value: val})
}

func (b *E) String(val string) *E {
	return b.Write(&Bytes{value: []byte(val)})
}

func (b *E) Bytes(val []byte) *E {
	return b.Write(&Bytes{value: val})
}

func (b *E) Strings(values []string) *E {
	n := len(values)
	b.Uint64(uint64(n))
	for _, val := range values {
		b.String(val)
	}
	return b
}

func (b *E) Encode() ([]byte, error) {
	buff := bytes.Buffer{}

	for _, vt := range b.values {
		raw, err := vt.Encode()
		if err != nil {
			return nil, err
		}
		_, err = buff.Write(raw)
		if err != nil {
			return nil, fmt.Errorf("can't write to buffer: %w", err)
		}
	}

	return buff.Bytes(), nil
}
