package cfg

import (
	"flag"
	"os"
	"strings"
	"sync"

	"github.com/injoyai/conv"
)

var (
	defaultFlags = &_flags{}
	onceFlags    sync.Once
)

func WithFlag[T Flag | *Flag | string](flags ...T) conv.IGetVar {
	onceFlags.Do(func() {
		f := &_flags{FlagSet: flag.NewFlagSet(os.Args[0], flag.ContinueOnError)}
		for _, v := range flags {
			switch vv := any(v).(type) {
			case *Flag:
				f.String(vv.Name, conv.String(vv.Default), vv.Usage)
			case Flag:
				f.String(vv.Name, conv.String(vv.Default), vv.Usage)
			case string:
				vv = strings.TrimLeft(vv, "-")
				f.String(vv, "", "")
			}
		}
		f.Parse(os.Args[1:])
		defaultFlags = f
	})
	return defaultFlags
}

type _flags struct {
	*flag.FlagSet
}

func (this *_flags) GetVar(key string) *conv.Var {
	f := this.Lookup(key)
	if f == nil || f.Value.String() == "" {
		return conv.Nil()
	}
	return conv.New(f.Value.String())
}

type Flag struct {
	Name    string //名称
	Default any    //默认值
	Usage   string //使用说明
}
