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

func (b *E) WInt(val int) *E {
	return b.Write(&Uint64{value: uint64(val)})
}

func (b *E) WInt64(val int64) *E {
	return b.Write(&Uint64{value: uint64(val)})
}

func (b *E) WUint64(val uint64) *E {
	return b.Write(&Uint64{value: val})
}

func (b *E) WString(val string) *E {
	return b.Write(&Bytes{value: []byte(val)})
}

func (b *E) WBytes(val []byte) *E {
	return b.Write(&Bytes{value: val})
}

func (b *E) WStrings(values []string) *E {
	n := len(values)
	b.WInt(n)
	for _, val := range values {
		b.WString(val)
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
