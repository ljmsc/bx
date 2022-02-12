package bx

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/matryer/is"
)

type Dummy struct {
	Name  string
	Value int32
}

func (d *Dummy) Encode() ([]byte, error) {
	return Encoder().String(d.Name).Int32(d.Value).Encode()
}

func (d *Dummy) Decode(_data []byte) (int, error) {
	startLen := len(_data)
	dec := Decoder(_data)
	var err error
	d.Name, err = dec.String()
	if err != nil {
		return 0, err
	}
	d.Value, err = dec.Int32()
	if err != nil {
		return 0, err
	}
	return startLen - dec.Size(), nil
}

func TestEncoder(t *testing.T) {
	is := is.New(t)

	valInt := int32(13)
	valUint64 := uint64(rand.Int63()) // #nosec
	valStr := "my text"

	raw, err := Encoder().Int32(valInt).String(valStr).Uint64(valUint64).Encode()
	is.NoErr(err)
	is.Equal(len(raw), 4+8+len(valStr)+8)
}

func TestEncodeStrings(t *testing.T) {
	is := is.New(t)
	testVals := []string{"val1", "val2", "val3"}
	size := len(strings.Join(testVals, ""))

	raw, err := Encoder().Strings(testVals).Encode()
	is.NoErr(err)
	is.Equal(len(raw), 4*8+size)
}

func TestEncodeFloa32(t *testing.T) {
	is := is.New(t)
	testVal := float32(1.56)
	raw, err := Encoder().Float32(testVal).Encode()
	is.NoErr(err)
	is.Equal(4, len(raw))
}

func TestEncodeFloa64(t *testing.T) {
	is := is.New(t)
	testVal := float64(1.56)
	raw, err := Encoder().Float64(testVal).Encode()
	is.NoErr(err)
	is.Equal(8, len(raw))
}

func TestEncodeTime(t *testing.T) {
	is := is.New(t)
	testVal := time.Date(2020, 03, 01, 12, 00, 45, 0, time.UTC)
	raw, err := Encoder().Time(testVal).Encode()
	is.NoErr(err)
	is.Equal(8, len(raw))
}

func TestEncodeDecodeValueTypeSlice(t *testing.T) {
	is := is.New(t)
	testData := []Dummy{
		{
			Name:  "test1",
			Value: 1,
		},
		{
			Name:  "test2",
			Value: 2,
		},
		{
			Name:  "test3",
			Value: 3,
		},
	}
	testDataVt := make([]ValueType, 0, len(testData))
	for _, td := range testData {
		ltd := td
		testDataVt = append(testDataVt, &ltd)
	}
	raw, err := Encoder().WriteSlice(testDataVt).Encode()
	is.NoErr(err)

	dec := Decoder(raw)
	resultVt, err := dec.ReadSlice(func() ValueType {
		return &Dummy{}
	})
	is.NoErr(err)
	is.Equal(3, len(resultVt))

	result := make([]Dummy, 0, len(resultVt))
	for _, vt := range resultVt {
		result = append(result, *vt.(*Dummy))
	}
	is.Equal(testData, result)
}
