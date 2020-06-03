package lbd

import "encoding/json"

type ServiceInformation struct {
	ServiceID   string `json:"serviceId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

func (l LBD) RetrieveServiceInformation(serviceId string) (*ServiceInformation, error) {
	r := NewGetRequest("/v1/services/" + serviceId)

	data, err := l.Do(r, nil, true)
	if err != nil {
		return nil, err
	}

	ret := new(ServiceInformation)
	return ret, json.Unmarshal(data.ResponseData, ret)
}
