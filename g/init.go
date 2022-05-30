package g

import (
	"errors"
	"github.com/injoyai/conv"
	"github.com/injoyai/conv/cfg"
)

type (
	Any interface{}
	Map map[string]interface{}
)

func (this Map) String() string {
	return conv.String(this)
}

func Cfg() *cfg.Entity {
	return cfg.Default
}

func NewCfg(paths ...string) *cfg.Entity {
	return cfg.New(paths...)
}

func NewVar(i interface{}) *conv.Var {
	return conv.New(i)
}

func NewMap(i interface{}) Map {
	return conv.New(i).Map()
}

func Recover(err *error) {
	if e := recover(); e != nil {
		*err = errors.New(conv.String(e))
	}
}
