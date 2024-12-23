package yaml

import "gopkg.in/yaml.v3"

type Yaml struct{}

func (Yaml) Marshal(v interface{}) ([]byte, error) {
	switch val := v.(type) {
	case []byte:
		return val, nil
	case string:
		return []byte(val), nil
	}
	return yaml.Marshal(v)
}

func (Yaml) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
