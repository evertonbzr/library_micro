package utils

import "encoding/json"

func DecodeJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func EncodeJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
