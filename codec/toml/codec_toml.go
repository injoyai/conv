package toml

import "github.com/pelletier/go-toml/v2"

type Toml struct{}

func (Toml) Marshal(v any) ([]byte, error) {
	return toml.Marshal(v)
}

func (Toml) Unmarshal(data []byte, v any) error {
	return toml.Unmarshal(data, v)
}
