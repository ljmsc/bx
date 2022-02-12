package bx

import (
	"bytes"
	"fmt"
	"time"
)

// E is the buffer to encode collected values to one byte slice
type E struct {
	values []ValueType
}

// Encoder creates a new encoder
func Encoder() *E {
	e := E{
		values: []ValueType{},
	}
	return &e
}

// Write writes any value which implements ValueType to the encoder
func (e *E) Write(vt ValueType) *E {
	if e.values == nil {
		e.values = []ValueType{}
	}
	e.values = append(e.values, vt)
	return e
}

// WriteSlice writes any slice of value which implements ValueType
func (e *E) WriteSlice(vt []ValueType) *E {
	if e.values == nil {
		e.values = []ValueType{}
	}
	e.Int64(int64(len(vt)))
	e.values = append(e.values, vt...)
	return e
}

// Int8 writes a int8 value to the encoder
func (e *E) Int8(val int8) *E {
	return e.Write(&Number{value: val})
}

// Int16 writes a int16 value to the encoder
func (e *E) Int16(val int16) *E {
	return e.Write(&Number{value: val})
}

// Int32 writes a int32 value to the encoder
func (e *E) Int32(val int32) *E {
	return e.Write(&Number{value: val})
}

// Int64 writes a int64 value to the encoder
func (e *E) Int64(val int64) *E {
	return e.Write(&Number{value: val})
}

// Uint8 writes a uint8 value to the encoder
func (e *E) Uint8(val uint8) *E {
	return e.Write(&Number{value: val})
}

// Uint16 writes a uint16 value to the encoder
func (e *E) Uint16(val uint16) *E {
	return e.Write(&Number{value: val})
}

// Uint32 writes a uint32 value to the encoder
func (e *E) Uint32(val uint32) *E {
	return e.Write(&Number{value: val})
}

// Uint64 writes a uint64 value to the encoder
func (e *E) Uint64(val uint64) *E {
	return e.Write(&Number{value: val})
}

// Float32 writes a float32 value to the encoder
func (e *E) Float32(val float32) *E {
	return e.Write(&Number{value: val})
}

// Float64 writes a float64 value to the encoder
func (e *E) Float64(val float64) *E {
	return e.Write(&Number{value: val})
}

// String writes a string value to the encoder
func (e *E) String(val string) *E {
	return e.Write(&Bytes{value: []byte(val)})
}

// Bytes writes a byte slice value to the encoder
func (e *E) Bytes(val []byte) *E {
	return e.Write(&Bytes{value: val})
}

// Strings writes a string slice value to the encoder
func (e *E) Strings(values []string) *E {
	n := len(values)
	e.Uint64(uint64(n))
	for _, val := range values {
		e.String(val)
	}
	return e
}

// Time writes a time value to the encoder
func (e *E) Time(val time.Time) *E {
	return e.Int64(val.UnixNano())
}

// Encode encodes all written values
func (e *E) Encode() ([]byte, error) {
	buff := bytes.Buffer{}

	for _, vt := range e.values {
		raw, err := vt.Encode()
		if err != nil {
			return nil, fmt.Errorf("can encode value type: %w", err)
		}
		_, err = buff.Write(raw)
		if err != nil {
			return nil, fmt.Errorf("can't write to buffer: %w", err)
		}
	}

	return buff.Bytes(), nil
}
