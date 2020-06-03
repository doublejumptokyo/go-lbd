package lbd

func (l LBD) RetrieveServerTime() (int64, error) {
	r := NewGetRequest("/v1/time")
	ret, err := l.Do(r, nil, false)
	if err != nil {
		return 0, err
	}
	return ret.ResponseTime, nil
}
