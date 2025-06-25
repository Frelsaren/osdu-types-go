package workproduct

type FluffyCoordinate struct {
	Double      *float64
	DoubleArray []float64
}

func (x *FluffyCoordinate) UnmarshalJSON(data []byte) error {
	x.DoubleArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.DoubleArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *FluffyCoordinate) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.DoubleArray != nil, x.DoubleArray, false, nil, false, nil, false, nil, false)
}
