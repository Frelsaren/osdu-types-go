package referencedata

type PurpleCoordinate struct {
	Double     *float64
	UnionArray []FluffyCoordinate
}

func (x *PurpleCoordinate) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.UnionArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *PurpleCoordinate) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.UnionArray != nil, x.UnionArray, false, nil, false, nil, false, nil, false)
}
