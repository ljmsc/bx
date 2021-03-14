package bx

import (
	"testing"

	"github.com/matryer/is"
)

func TestDecoder(t *testing.T) {
	is := is.New(t)

	val1 := 13
	val2 := "my text"
	raw, _ := Encoder().WInt(val1).WString(val2).Encode()

	dec := Decoder(raw)
	rval1, err := dec.RInt()
	is.NoErr(err)
	is.Equal(rval1, val1)

	rval2, err := dec.RString()
	is.NoErr(err)
	is.Equal(rval2, val2)
}

func TestWStrings(t *testing.T) {
	is := is.New(t)
	testVals := []string{"val1", "val2", "val3"}
	raw, _ := Encoder().WStrings(testVals).Encode()

	dec := Decoder(raw)
	retTestVals, err := dec.RStrings()
	is.NoErr(err)
	is.Equal(retTestVals, testVals)
}
