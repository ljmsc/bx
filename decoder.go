package bx

type D struct {
	raw []byte
}

func Decoder(raw []byte) *D {
	d := D{
		raw: []byte{},
	}
	d.raw = raw
	return &d
}

func (d *D) Read(vt ValueType) error {
	n, err := vt.Decode(d.raw)
	if err != nil {
		return err
	}
	d.raw = d.raw[n:]
	return nil
}

func (d *D) Int8() (int8, error) {
	var val int8
	vt := Number{size: 1, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

func (d *D) Int16() (int16, error) {
	var val int16
	vt := Number{size: 2, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

func (d *D) Int32() (int32, error) {
	var val int32
	vt := Number{size: 4, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

func (d *D) Int64() (int64, error) {
	var val int64
	vt := Number{size: 8, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

func (d *D) Uint8() (uint8, error) {
	var val uint8
	vt := Number{size: 1, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

func (d *D) Uint16() (uint16, error) {
	var val uint16
	vt := Number{size: 2, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

func (d *D) Uint32() (uint32, error) {
	var val uint32
	vt := Number{size: 4, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

func (d *D) Uint64() (uint64, error) {
	var val uint64
	vt := Number{size: 8, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

func (d *D) Float32() (float32, error) {
	var val float32
	vt := Number{size: 4, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

func (d *D) Float64() (float64, error) {
	var val float64
	vt := Number{size: 8, value: &val}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return val, nil
}

func (d *D) String() (string, error) {
	vt := Bytes{}
	if err := d.Read(&vt); err != nil {
		return "", err
	}
	return vt.String(), nil
}

func (d *D) Bytes() ([]byte, error) {
	vt := Bytes{}
	if err := d.Read(&vt); err != nil {
		return nil, err
	}
	return vt.Value(), nil
}

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
