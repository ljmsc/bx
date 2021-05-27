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

func (e *E) Write(vt ValueType) *E {
	if e.values == nil {
		e.values = []ValueType{}
	}
	e.values = append(e.values, vt)
	return e
}

func (e *E) Int8(val int8) *E {
	return e.Write(&Number{value: val})
}

func (e *E) Int16(val int16) *E {
	return e.Write(&Number{value: val})
}

func (e *E) Int32(val int32) *E {
	return e.Write(&Number{value: val})
}

func (e *E) Int64(val int64) *E {
	return e.Write(&Number{value: val})
}

func (e *E) Uint8(val uint8) *E {
	return e.Write(&Number{value: val})
}

func (e *E) Uint16(val uint16) *E {
	return e.Write(&Number{value: val})
}

func (e *E) Uint32(val uint32) *E {
	return e.Write(&Number{value: val})
}

func (e *E) Uint64(val uint64) *E {
	return e.Write(&Number{value: val})
}

func (e *E) Float32(val float32) *E {
	return e.Write(&Number{value: val})
}

func (e *E) Float64(val float64) *E {
	return e.Write(&Number{value: val})
}

func (e *E) String(val string) *E {
	return e.Write(&Bytes{value: []byte(val)})
}

func (e *E) Bytes(val []byte) *E {
	return e.Write(&Bytes{value: val})
}

func (e *E) Strings(values []string) *E {
	n := len(values)
	e.Uint64(uint64(n))
	for _, val := range values {
		e.String(val)
	}
	return e
}

func (e *E) Encode() ([]byte, error) {
	buff := bytes.Buffer{}

	for _, vt := range e.values {
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
