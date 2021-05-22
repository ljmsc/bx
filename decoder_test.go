package bx

import (
	"testing"

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

func TestWStrings(t *testing.T) {
	is := is.New(t)
	testVals := []string{"val1", "val2", "val3"}
	raw, _ := Encoder().Strings(testVals).Encode()

	dec := Decoder(raw)
	retTestVals, err := dec.Strings()
	is.NoErr(err)
	is.Equal(retTestVals, testVals)
}
