package bx

import (
	"bytes"
	"encoding/binary"
)

const (
	uint64Size = 8
)

// ValueType is the basic interface for encoding and decoding of values
type ValueType interface {
	Decode(_data []byte) (int, error)
	Encode() ([]byte, error)
}

// Bytes is the encoder / decoder for byte slices
type Bytes struct {
	value []byte
}

// String returns the string representation of the byte slice
func (b *Bytes) String() string {
	return string(b.value)
}

// Decode decodes the given data to a byte slice
func (b *Bytes) Decode(_data []byte) (int, error) {
	if len(_data) < 8 {
		return 0, ErrNotEnoughData
	}
	var val uint64
	sizevt := &Number{size: 8, value: &val}
	if _, err := sizevt.Decode(_data); err != nil {
		return 0, err
	}
	size := int(val)
	_data = _data[uint64Size:]
	if len(_data) < size {
		return 0, ErrInvalidData
	}
	b.value = _data[:size]

	return size + uint64Size, nil
}

// Encode encodes a byte slice
func (b *Bytes) Encode() ([]byte, error) {
	rawSize, _ := (&Number{value: uint64(len(b.value))}).Encode()
	raw := make([]byte, 0, len(rawSize)+len(b.value))
	raw = append(raw, rawSize...)
	raw = append(raw, b.value...)
	return raw, nil
}

// Number is the encoder / decoder for number values
type Number struct {
	value interface{}
	size  int
}

// Decode decodes the given data to a number value
func (n *Number) Decode(_data []byte) (int, error) {
	if len(_data) < n.size {
		return 0, ErrNotEnoughData
	}
	buff := bytes.NewBuffer(_data[:n.size])
	if err := binary.Read(buff, binary.LittleEndian, n.value); err != nil {
		return 0, err
	}
	return n.size, nil
}

// Encode encodes a number value
func (n *Number) Encode() ([]byte, error) {
	buff := bytes.Buffer{}
	if err := binary.Write(&buff, binary.LittleEndian, n.value); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
