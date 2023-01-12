package lbd

// Deprecated: Use ConstructRawTransactionV2
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

func ConstructRawTransactionV2(raw []byte) (*TransactionV2, error) {
	ret, err := UnmarshalTransactionV2(raw)
	if err != nil {
		return nil, err
	}
	err = ret.Check()
	if err != nil {
		return nil, err
	}
	return ret, nil
}
