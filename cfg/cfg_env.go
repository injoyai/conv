package cfg

import (
	"github.com/injoyai/conv/v2"
	"os"
)

var defaultEnv = &Env{}

func WithEnv() conv.StringGetter {
	return defaultEnv
}

type Env struct{}

func (this *Env) GetVar(key string) *conv.Var {
	if v, ok := os.LookupEnv(key); ok {
		return conv.New(v)
	}
	return conv.Nil
}
