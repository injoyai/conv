package cfg

import (
	"github.com/injoyai/conv"
	"os"
)

var defaultEnv = &Env{}

func WithEnv() conv.IGetVar {
	return defaultEnv
}

type Env struct{}

func (this *Env) GetVar(key string) *conv.Var {
	if v, ok := os.LookupEnv(key); ok {
		return conv.New(v)
	}
	return conv.Nil()
}
