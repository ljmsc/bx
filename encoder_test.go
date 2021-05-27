package bx

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestEncoder(t *testing.T) {
	is := is.New(t)

	valInt := int32(13)
	valUint64 := uint64(rand.Int63())
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
