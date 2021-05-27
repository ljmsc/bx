# bx
[![Go Report Card](https://goreportcard.com/badge/github.com/ljmsc/bx)](https://goreportcard.com/report/github.com/ljmsc/bx)
[![GoDoc](https://godoc.org/github.com/ljmsc/bx?status.svg)](https://pkg.go.dev/mod/github.com/ljmsc/bx?tab=overview)

bx is a binary encoding/decoding lib for go.

## Installation
The code was tested with go version `1.15`.
Older versions can also work but have not been tested.
```
go get -u github.com/ljmsc/bx
```

## Usage

### encoding
```go
package main

import (
	"github.com/ljmsc/bx"
)

func main() {
	foo := "this is my test value string"
	bar := int64(1337)

	raw, err := bx.Encoder().String(foo).Int64(bar).Encode()
	if err != nil {
		// handle error
	}
	// do something with the encoded data
}
```

### decoding
```go
package main

import (
	"github.com/ljmsc/bx"
)

func main() {
	var raw []byte // some binary data

	dec := bx.Decoder(raw)
	foo, err := dec.String()
	if err != nil {
		// handle error
	}
	bar, err := dec.Int64()
	if err != nil {
		// handle error
	}

	// use foo and bar
}
```


## License
The project (and all code) is licensed under the Mozilla Public License Version 2.0.
