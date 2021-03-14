package bx

import (
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
	sizevt := &Uint64{}
	if _, err := sizevt.Decode(_data); err != nil {
		return 0, err
	}
	size := int(sizevt.Value())
	_data = _data[uint64Size:]
	if len(_data) < size {
		return 0, ErrInvalidData
	}
	b.value = _data[:size]

	return size + uint64Size, nil
}

func (b *Bytes) Encode() ([]byte, error) {
	rawSize, _ := (&Uint64{value: uint64(len(b.value))}).Encode()
	raw := make([]byte, 0, len(rawSize)+len(b.value))
	raw = append(raw, rawSize...)
	raw = append(raw, b.value...)
	return raw, nil
}

type Uint64 struct {
	value uint64
}

func (u *Uint64) Value() uint64 {
	return u.value
}

func (u *Uint64) Decode(_data []byte) (int, error) {
	if len(_data) < uint64Size {
		return 0, ErrInsufficientData
	}
	_data = _data[:uint64Size]
	u.value = binary.LittleEndian.Uint64(_data)
	return uint64Size, nil
}

func (u *Uint64) Encode() ([]byte, error) {
	raw := make([]byte, uint64Size)
	binary.LittleEndian.PutUint64(raw, u.value)
	return raw, nil
}
