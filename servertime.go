package lbd

func (l LBD) RetrieveServerTime() (int64, error) {
	path := "/v1/time"
	ret, err := l.get(path, "", false)
	if err != nil {
		return 0, err
	}
	return ret.ResponseTime, nil
}
