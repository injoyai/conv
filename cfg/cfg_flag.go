package cfg

import (
	"flag"
	"github.com/injoyai/conv/v2"
	"os"
	"sync"
)

var (
	defaultFlags = &Flags{}
	onceFlags    sync.Once
)

func WithFlag(flags ...*Flag) conv.StringGetter {
	onceFlags.Do(func() {
		f := &Flags{FlagSet: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
		for _, v := range flags {
			f.String(v.Name, conv.String(v.Default), v.Usage)
		}
		f.Parse(os.Args[1:])
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
		return conv.Nil
	}
	return conv.New(f.Value.String())
}

type Flag struct {
	Name    string //名称
	Default any    //默认值
	Usage   string //使用说明
}
