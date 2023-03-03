package helpers

import "encoding/json"

func GetRequestBody(obj interface{}) (string, error) {
	var res string
	marshaled, err := json.Marshal(obj)
	if err != nil {
		return res, err
	}

	res = string(marshaled)
	return res, nil
}
