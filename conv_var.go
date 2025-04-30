package conv

import (
	"errors"
	"time"
)

var (
	Nil = &Var{}
)

func New(i any) *Var {
	if val, ok := i.(*Var); ok {
		return val
	}
	return &Var{canSet: true, value: i}
}

type Var struct {
	canSet bool //是否能改变值
	value  any  //值
}

func (this *Var) Val() any {
	return this.value
}

func (this *Var) Set(i any) {
	if !this.canSet {
		return
	}
	this.value = i
}

func (this *Var) IsDefault() bool {
	return IsDefault(this.Val())
}

func (this *Var) IsNumber() bool {
	return IsNumber(this.Val())
}

func (this *Var) IsString() bool {
	return IsString(this.Val())
}

func (this *Var) IsInt() bool {
	return IsInt(this.Val())
}

func (this *Var) IsFloat() bool {
	return IsFloat(this.Val())
}

func (this *Var) IsBool() bool {
	return IsBool(this.Val())
}

func (this *Var) IsTime() bool {
	return IsTime(this.Val())
}

func (this *Var) IsArray() bool {
	return IsArray(this.Val())
}

func (this *Var) IsPointer() bool {
	return IsPointer(this.Val())
}

func (this *Var) IsNil() bool {
	return this == nil || this.Val() == nil
}

func (this *Var) Interface(def ...any) any {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return this.Val()
}

func (this *Var) Interfaces(def ...[]any) []any {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Interfaces(this.Val())
}

func (this *Var) Byte(def ...byte) byte {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Byte(this.Val())
}

func (this *Var) Bytes(def ...[]byte) []byte {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Bytes(this.Val())
}

func (this *Var) Rune(def ...rune) rune {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Rune(this.Val())
}

func (this *Var) Runes(def ...[]rune) []rune {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Runes(this.Val())
}

func (this *Var) String(def ...string) string {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return String(this.Val())
}

func (this *Var) Strings(def ...[]string) []string {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Strings(this.Val())
}

func (this *Var) Int(def ...int) int {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int(this.Val())
}

func (this *Var) Ints(def ...[]int) []int {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Ints(this.Val())
}

func (this *Var) Int8(def ...int8) int8 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int8(this.Val())
}

func (this *Var) Int16(def ...int16) int16 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int16(this.Val())
}

func (this *Var) Int32(def ...int32) int32 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int32(this.Val())
}

func (this *Var) Int64(def ...int64) int64 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int64(this.Val())
}

func (this *Var) Int64s(def ...[]int64) []int64 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int64s(this.Val())
}

func (this *Var) Uint(def ...uint) uint {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Uint(this.Val())
}

func (this *Var) Uint8(def ...uint8) uint8 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Uint8(this.Val())
}

func (this *Var) Uint16(def ...uint16) uint16 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Uint16(this.Val())
}

func (this *Var) Uint32(def ...uint32) uint32 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Uint32(this.Val())
}

func (this *Var) Uint64(def ...uint64) uint64 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Uint64(this.Val())
}

func (this *Var) Float32(def ...float32) float32 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Float32(this.Val())
}

func (this *Var) Float64(def ...float64) float64 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Float64(this.Val())
}

func (this *Var) Bool(def ...bool) bool {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Bool(this.Val())
}

func (this *Var) Duration(def ...time.Duration) time.Duration {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return time.Duration(this.Int64())
}

func (this *Var) Microsecond(def ...time.Duration) time.Duration {
	return this.Duration(def...) * time.Microsecond
}

func (this *Var) Millisecond(def ...time.Duration) time.Duration {
	return this.Duration(def...) * time.Millisecond
}

func (this *Var) Second(def ...time.Duration) time.Duration {
	return this.Duration(def...) * time.Second
}

func (this *Var) Minute(def ...time.Duration) time.Duration {
	return this.Duration(def...) * time.Minute
}

func (this *Var) Hour(def ...time.Duration) time.Duration {
	return this.Duration(def...) * time.Hour
}

func (this *Var) Err(def ...error) error {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	if this.IsNil() {
		return nil
	}
	return errors.New(this.String())
}

func (this *Var) Map(def ...map[string]any) map[string]any {
	return this.GMap(def...)
}

func (this *Var) GMap(def ...map[string]any) map[string]any {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return GMap(this.Val())
}

func (this *Var) SMap(def ...map[string]string) map[string]string {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return SMap(this.Val())
}

func (this *Var) IMap(def ...map[any]any) map[any]any {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return IMap(this.Val())
}

func (this *Var) AMap(def ...map[any]any) map[any]any {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return IMap(this.Val())
}

func (this *Var) DMap(def ...any) *Map {
	if this.IsNil() && len(def) > 0 {
		return NewMap(def[0])
	}
	return NewMap(this.Val())
}

func (this *Var) Copy() *Var {
	return New(Copy(this.Val()))
}

func (this *Var) Unmarshal(ptr any) error {
	return Unmarshal(this.Val(), ptr)
}
