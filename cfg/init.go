package cfg

import (
	"github.com/injoyai/conv"
	"io/ioutil"
	"time"
)

var (
	defaultPath = "./config/config.json"
	Default     = New()
)

type Entity struct {
	err  error
	path string
	conv.Map
}

func (this *Entity) Reread() *Entity {
	x := New(this.path)
	if x.err == nil {
		this.Map = x.Map
	}
	return this
}

func (this *Entity) String() string {
	if this.err != nil {
		return this.err.Error()
	}
	return this.Map.Var.String()
}

func (this *Entity) Error() error {
	return this.err
}

func New(paths ...string) *Entity {
	path := conv.GetDefaultString(defaultPath, paths...)
	data := &Entity{path: path}
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		data.err = err
	}
	data.Map = conv.NewMap(bs)
	return data
}

// Reread 重新读取配置
func Reread() *Entity {
	Default.Reread()
	return Default
}

// Get 获取value类型数据
func Get(key string) conv.Map {
	return Default.Get(key)
}

// GetVar 获取var类型数据
func GetVar(key string) *conv.Var {
	return Default.GetVar(key)
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
