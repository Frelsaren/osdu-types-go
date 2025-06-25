package workproduct

type AnyCRSGeoJSONPointCoordinate struct {
	Double     *float64
	UnionArray []PurpleCoordinate
}

func (x *AnyCRSGeoJSONPointCoordinate) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.UnionArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AnyCRSGeoJSONPointCoordinate) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.UnionArray != nil, x.UnionArray, false, nil, false, nil, false, nil, false)
}
