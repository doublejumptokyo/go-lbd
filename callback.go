package lbd

func ConstructRawTransaction(raw []byte) (*Transaction, error) {
	ret, err := UnmarshalTransaction(raw)
	if err != nil {
		return nil, err
	}

	err = ret.Check()
	if err != nil {
		return nil, err
	}

	return ret, nil
}
