package cfg

import (
	"github.com/injoyai/conv/v2"
	"github.com/injoyai/conv/v2/codec"
	"os"
	"path/filepath"
	"strings"
)

func WithAny(i any, codec ...codec.Interface) conv.StringGetter {
	return conv.NewMap(i, codec...)
}

func WithDefaultFile() conv.StringGetter {
	return WithFile("./config/config.yaml")
}

func WithExecuteFile(join string, codecs ...codec.Interface) conv.StringGetter {
	executeName, err := os.Executable()
	if err != nil {
		executeName = "./"
	}
	dir := filepath.Dir(executeName)
	filename := filepath.Join(dir, join)
	return WithFile(filename, codecs...)
}

func WithFile(filename string, codecs ...codec.Interface) conv.StringGetter {
	bs, err := os.ReadFile(filename)
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

func WithYaml(filename string) conv.StringGetter {
	return WithFile(filename, codec.Yaml)
}

func WithJson(filename string) conv.StringGetter {
	return WithFile(filename, codec.Json)
}

func WithToml(filename string) conv.StringGetter {
	return WithFile(filename, codec.Toml)
}

func WithIni(filename string) conv.StringGetter {
	return WithFile(filename, codec.Ini)
}
