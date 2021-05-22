package bx

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var (
	ErrInsufficientData = fmt.Errorf("insufficient data")
	ErrInvalidData      = fmt.Errorf("invalid data")
)

const (
	uint64Size = 8
)

type ValueType interface {
	Decode(_data []byte) (int, error)
	Encode() ([]byte, error)
}

type Bytes struct {
	value []byte
}

func (b *Bytes) String() string {
	return string(b.value)
}

func (b *Bytes) Value() []byte {
	return b.value
}

func (b *Bytes) Decode(_data []byte) (int, error) {
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

func (b *Bytes) Encode() ([]byte, error) {
	rawSize, _ := (&Number{value: uint64(len(b.value))}).Encode()
	raw := make([]byte, 0, len(rawSize)+len(b.value))
	raw = append(raw, rawSize...)
	raw = append(raw, b.value...)
	return raw, nil
}

type Number struct {
	value interface{}
	size  int
}

func (n *Number) Decode(_data []byte) (int, error) {
	buff := bytes.NewBuffer(_data[:n.size])
	if err := binary.Read(buff, binary.LittleEndian, n.value); err != nil {
		return 0, err
	}
	return n.size, nil
}

func (n *Number) Encode() ([]byte, error) {
	buff := bytes.Buffer{}
	if err := binary.Write(&buff, binary.LittleEndian, n.value); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
