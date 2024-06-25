package cfg

import (
	"github.com/injoyai/conv"
	"time"
)

var Default = New(WithDefaultFile())

func Init(i ...conv.IGetVar) { Default = New(i...) }

func New(i ...conv.IGetVar) *Entity {
	c := &Entity{}
	for _, v := range i {
		if v != nil {
			c.list = append(c.list, conv.NewExtend(v))
		}
	}
	c.Extend = conv.NewExtend(c)
	return c
}

type Entity struct {
	list []conv.Extend
	conv.Extend
}

func (this *Entity) GetVar(key string) *conv.Var {
	for _, v := range this.list {
		if val := v.GetVar(key); !val.IsNil() {
			return val
		}
	}
	return conv.Nil()
}

// GetString 读取string类型数据,其它类型可转string类型,复杂的类型会转json
func GetString(key string, def ...string) string {
	return Default.GetString(key, def...)
}

// GetStrings 读取字符串数组数据
func GetStrings(key string, def ...[]string) []string {
	return Default.GetStrings(key, def...)
}

// GetInt 读取int类型数据
func GetInt(key string, def ...int) int {
	return Default.GetInt(key, def...)
}

// GetInts 读取[]int类型数据
func GetInts(key string, def ...[]int) []int {
	return Default.GetInts(key, def...)
}

// GetInt64 读取int64类型数据
func GetInt64(key string, def ...int64) int64 {
	return Default.GetInt64(key, def...)
}

// GetInt64s 读取[]int64类型数据
func GetInt64s(key string, def ...[]int64) []int64 {
	return Default.GetInt64s(key, def...)
}

// GetFloat32 读取float32类型数据
func GetFloat32(key string, def ...float32) float32 {
	return Default.GetFloat32(key, def...)
}

// GetFloat64 读取float64类型数据
func GetFloat64(key string, def ...float64) float64 {
	return Default.GetFloat64(key, def...)
}

// GetBool 读取bool类型数据
func GetBool(key string, def ...bool) bool {
	return Default.GetBool(key, def...)
}

// GetDuration 读取时间类型,纳秒
func GetDuration(key string, def ...time.Duration) time.Duration {
	return Default.GetDuration(key, def...)
}

// GetSecond 读取时间类型,秒
func GetSecond(key string, def ...time.Duration) time.Duration {
	return Default.GetSecond(key, def...)
}

// GetMinute 读取时间类型,分
func GetMinute(key string, def ...time.Duration) time.Duration {
	return Default.GetMinute(key, def...)
}

// GetHour 读取时间类型,时
func GetHour(key string, def ...time.Duration) time.Duration {
	return Default.GetHour(key, def...)
}

// GetMap 读取map[string]interface{}类型
func GetMap(key string, def ...map[string]interface{}) map[string]interface{} {
	return Default.GetGMap(key, def...)
}

func GetDMap(key string, def ...interface{}) *conv.Map {
	return Default.GetDMap(key, def...)
}

func GetInterfaces(key string, def ...[]interface{}) []interface{} {
	return Default.GetInterfaces(key, def...)
}
