package cfg

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/conv/codec"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func WithAny(i interface{}, codec ...codec.Interface) conv.IGetVar {
	return conv.NewMap(i, codec...)
}

func WithDefaultFile() conv.IGetVar {
	return WithFile("./config/config.yaml")
}

func WithFile(filename string, codecs ...codec.Interface) conv.IGetVar {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//log.Println("err: ", err)
		//return nil, err
	}

	if len(codecs) == 0 {
		//根据文件类型
		switch strings.ToLower(filepath.Ext(filename)) {
		case ".yaml", ".yml":
			codecs = append(codecs, codec.Yaml)
		case ".ini":
			codecs = append(codecs, codec.Ini)
		case ".toml":
			codecs = append(codecs, codec.Toml)
		case ".json":
			codecs = append(codecs, codec.Json)
		}
	}

	return conv.NewMap(bs, codecs...)
}

func WithYaml(filename string) conv.IGetVar {
	return WithFile(filename, codec.Yaml)
}

func WithJson(filename string) conv.IGetVar {
	return WithFile(filename, codec.Json)
}

func WithToml(filename string) conv.IGetVar {
	return WithFile(filename, codec.Toml)
}

func WithIni(filename string) conv.IGetVar {
	return WithFile(filename, codec.Ini)
}
