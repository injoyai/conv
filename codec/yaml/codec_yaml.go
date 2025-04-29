package yaml

import "gopkg.in/yaml.v3"

type Yaml struct{}

func (Yaml) Marshal(v any) ([]byte, error) {
	switch val := v.(type) {
	case []byte:
		return val, nil
	case string:
		return []byte(val), nil
	}
	return yaml.Marshal(v)
}

func (Yaml) Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}
