package codec

import (
	"github.com/injoyai/conv/v2/codec/ini"
	"github.com/injoyai/conv/v2/codec/json"
	"github.com/injoyai/conv/v2/codec/toml"
	"github.com/injoyai/conv/v2/codec/yaml"
	"strings"
)

type Interface interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

var (
	Default Interface = Json
	Json    Interface = json.Json{}
	Toml    Interface = toml.Toml{}
	Yaml    Interface = yaml.Yaml{}
	Ini     Interface = ini.Ini{}
)

func Get(s string) Interface {
	s = strings.TrimLeft(s, ".")
	switch strings.ToLower(s) {
	case "json":
		return Json
	case "yaml", "yml":
		return Yaml
	case "toml":
		return Toml
	case "ini":
		return Ini
	default:
		return Default
	}
}
