package json

import json "github.com/json-iterator/go"

type Json struct{}

func (Json) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (Json) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
