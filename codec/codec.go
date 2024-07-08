package codec

import (
	"github.com/injoyai/conv/codec/ini"
	"github.com/injoyai/conv/codec/json"
	"github.com/injoyai/conv/codec/toml"
	"github.com/injoyai/conv/codec/yaml"
	"strings"
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
	Ini     Interface = ini.Ini{}
)

func Get(s string) Interface {
	switch strings.ToLower(s) {
	case "json":
		return Json
	case "yaml":
		return Yaml
	case "toml":
		return Toml
	case "ini":
		return Ini
	default:
		return Default
	}
}
