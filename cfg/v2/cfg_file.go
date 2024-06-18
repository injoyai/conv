package cfg

import (
	"github.com/injoyai/conv"
	"github.com/injoyai/conv/codec"
	"io/ioutil"
)

func WithAny(i interface{}, codec ...codec.Interface) conv.IGetVar {
	return conv.NewMap(i, codec...)
}

func WithFile(filename string, codec ...codec.Interface) conv.IGetVar {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//log.Println("err: ", err)
		//return nil, err
	}
	return conv.NewMap(bs, codec...)
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
