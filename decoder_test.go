package bx

import (
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestDecoder(t *testing.T) {
	is := is.New(t)

	val1 := int64(13)
	val2 := "my text"
	raw, _ := Encoder().Int64(val1).String(val2).Encode()

	dec := Decoder(raw)
	rval1, err := dec.Int64()
	is.NoErr(err)
	is.Equal(rval1, val1)

	rval2, err := dec.String()
	is.NoErr(err)
	is.Equal(rval2, val2)
}

func TestDecodeStrings(t *testing.T) {
	is := is.New(t)
	testVals := []string{"val1", "val2", "val3"}
	raw, _ := Encoder().Strings(testVals).Encode()

	dec := Decoder(raw)
	retTestVals, err := dec.Strings()
	is.NoErr(err)
	is.Equal(retTestVals, testVals)
}

func TestDecodeFloat32(t *testing.T) {
	is := is.New(t)
	testVal := float32(1.56)
	raw, _ := Encoder().Float32(testVal).Encode()

	dec := Decoder(raw)
	retTestVal, err := dec.Float32()
	is.NoErr(err)
	is.Equal(testVal, retTestVal)
}

func TestDecodeFloat64(t *testing.T) {
	is := is.New(t)
	testVal := float64(1.56)
	raw, _ := Encoder().Float64(testVal).Encode()

	dec := Decoder(raw)
	retTestVal, err := dec.Float64()
	is.NoErr(err)
	is.Equal(testVal, retTestVal)
}

func TestDecodeTime(t *testing.T) {
	is := is.New(t)
	testVal := time.Date(2020, 03, 01, 12, 00, 45, 0, time.UTC)
	raw, _ := Encoder().Time(testVal).Encode()

	dec := Decoder(raw)
	retTime, err := dec.Time()
	is.NoErr(err)
	is.Equal(testVal, retTime)
	is.Equal(testVal.Unix(), retTime.Unix())
}
