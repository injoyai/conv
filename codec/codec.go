package codec

import (
	"github.com/injoyai/conv/codec/json"
	"github.com/injoyai/conv/codec/toml"
	"github.com/injoyai/conv/codec/yaml"
)

type Interface interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

var (
	Default Interface = Json
	Json    Interface = json.Json{}
	Toml    Interface = toml.Toml{}
	Yaml    Interface = yaml.Yaml{}
)
