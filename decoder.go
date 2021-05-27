package bx

// D is the buffer for decoding values
type D struct {
	raw []byte
}

// Decoder creates a new decoder
func Decoder(raw []byte) *D {
	d := D{
		raw: []byte{},
	}
	d.raw = raw
	return &d
}

// Read reads any value which implements ValueType to the decoder
func (d *D) Read(vt ValueType) error {
	n, err := vt.Decode(d.raw)
	if err != nil {
		return err
	}
	d.raw = d.raw[n:]
	return nil
}

// Int8 reads a int8 value from the decoder
func (d *D) Int8() (int8, error) {
	var val int8
	vt := Number{size: 1, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

// Int16 reads a int16 value from the decoder
func (d *D) Int16() (int16, error) {
	var val int16
	vt := Number{size: 2, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

// Int32 reads a int32 value from the decoder
func (d *D) Int32() (int32, error) {
	var val int32
	vt := Number{size: 4, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

// Int64 reads a int64 value from the decoder
func (d *D) Int64() (int64, error) {
	var val int64
	vt := Number{size: 8, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

// Uint8 reads a uint8 value from the decoder
func (d *D) Uint8() (uint8, error) {
	var val uint8
	vt := Number{size: 1, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

// Uint16 reads a uint16 value from the decoder
func (d *D) Uint16() (uint16, error) {
	var val uint16
	vt := Number{size: 2, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

// Uint32 reads a uint32 value from the decoder
func (d *D) Uint32() (uint32, error) {
	var val uint32
	vt := Number{size: 4, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

// Uint64 reads a uint64 value from the decoder
func (d *D) Uint64() (uint64, error) {
	var val uint64
	vt := Number{size: 8, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

// Float32 reads a float32 value from the decoder
func (d *D) Float32() (float32, error) {
	var val float32
	vt := Number{size: 4, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

// Float64 reads a float64 value from the decoder
func (d *D) Float64() (float64, error) {
	var val float64
	vt := Number{size: 8, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

// String reads a string value from the decoder
func (d *D) String() (string, error) {
	vt := Bytes{}
	if err := d.Read(&vt); err != nil {
		return "", err
	}
	return vt.String(), nil
}

// Bytes reads a byte slice value from the decoder
func (d *D) Bytes() ([]byte, error) {
	vt := Bytes{}
	if err := d.Read(&vt); err != nil {
		return nil, err
	}
	return vt.value, nil
}

// Strings reads a string slice value from the decoder
func (d *D) Strings() ([]string, error) {
	n, err := d.Uint64()
	if err != nil {
		return nil, err
	}
	values := make([]string, 0, n)
	for i := uint64(0); i < n; i++ {
		val, err := d.String()
		if err != nil {
			return nil, err
		}
		values = append(values, val)
	}
	return values, nil
}
