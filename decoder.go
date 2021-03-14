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

func (d *D) RInt() (int, error) {
	val, err := d.RUint64()
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

func (d *D) RInt64() (int64, error) {
	val, err := d.RUint64()
	if err != nil {
		return 0, err
	}
	return int64(val), nil
}

func (d *D) RUint64() (uint64, error) {
	vt := Uint64{}
	if err := d.Read(&vt); err != nil {
		return 0, err
	}
	return vt.Value(), nil
}

func (d *D) RString() (string, error) {
	vt := Bytes{}
	if err := d.Read(&vt); err != nil {
		return "", err
	}
	return vt.String(), nil
}

func (d *D) RBytes() ([]byte, error) {
	vt := Bytes{}
	if err := d.Read(&vt); err != nil {
		return nil, err
	}
	return vt.Value(), nil
}

func (d *D) RStrings() ([]string, error) {
	n, err := d.RInt()
	if err != nil {
		return nil, err
	}
	values := make([]string, 0, n)
	for i := 0; i < n; i++ {
		val, err := d.RString()
		if err != nil {
			return nil, err
		}
		values = append(values, val)
	}
	return values, nil
}
