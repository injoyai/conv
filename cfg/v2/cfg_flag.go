package cfg

import (
	"flag"
	"github.com/injoyai/conv"
	"os"
	"strings"
	"sync"
)

var (
	defaultFlags = &Flags{}
	onceFlags    sync.Once
)

func WithFlag(flags ...*Flag) conv.IGetVar {
	onceFlags.Do(func() {
		f := &Flags{FlagSet: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
		for _, v := range flags {
			f.String(v.Name, conv.String(v.Default), v.Usage)
		}
		args := []string(nil)
		n := 1
		for i := range os.Args {
			if strings.HasPrefix(os.Args[i], "-") || i-n == 1 {
				args = append(args, os.Args[i])
				n = i
			}
		}
		f.Parse(args)
		defaultFlags = f
	})
	return defaultFlags
}

type Flags struct {
	*flag.FlagSet
}

func (this *Flags) GetVar(key string) *conv.Var {
	f := this.Lookup(key)
	if f == nil || f.Value.String() == "" {
		return conv.Nil()
	}
	return conv.New(f.Value.String())
}

type Flag struct {
	Name    string      //名称
	Default interface{} //默认值
	Usage   string      //使用说明
}
