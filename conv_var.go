package conv

import (
	"errors"
	"time"
)

type Var struct {
	Value interface{}
}

func Nil() *Var {
	return New(nil)
}

func New(i interface{}) *Var {
	return &Var{Value: i}
}

func (this *Var) IsDefault() bool {
	return IsDefault(this.Value)
}

func (this *Var) IsNumber() bool {
	return IsNumber(this.Value)
}

func (this *Var) IsInt() bool {
	return IsInt(this.Value)
}

func (this *Var) IsFloat() bool {
	return IsFloat(this.Value)
}

func (this *Var) IsBool() bool {
	return IsBool(this.Value)
}

func (this *Var) IsTime() bool {
	return IsTime(this.Value)
}

func (this *Var) IsNil() bool {
	return this.Value == nil
}

func (this *Var) Val() interface{} {
	return this.Value
}

func (this *Var) Interface(def ...interface{}) interface{} {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return this.Value
}

func (this *Var) Interfaces(def ...[]interface{}) []interface{} {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Interfaces(this.Value)
}

func (this *Var) Byte(def ...byte) byte {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Byte(this.Value)
}

func (this *Var) Bytes(def ...[]byte) []byte {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Bytes(this.Value)
}

func (this *Var) Rune(def ...rune) rune {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Rune(this.Value)
}

func (this *Var) Runes(def ...[]rune) []rune {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Runes(this.Value)
}

func (this *Var) String(def ...string) string {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return String(this.Value)
}

func (this *Var) Strings(def ...[]string) []string {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Strings(this.Value)
}

func (this *Var) Int(def ...int) int {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int(this.Value)
}

func (this *Var) Ints(def ...[]int) []int {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Ints(this.Value)
}

func (this *Var) Int8(def ...int8) int8 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int8(this.Value)
}

func (this *Var) Int16(def ...int16) int16 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int16(this.Value)
}

func (this *Var) Int32(def ...int32) int32 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int32(this.Value)
}

func (this *Var) Int64(def ...int64) int64 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int64(this.Value)
}

func (this *Var) Int64s(def ...[]int64) []int64 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Int64s(this.Value)
}

func (this *Var) Uint(def ...uint) uint {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Uint(this.Value)
}

func (this *Var) Uint8(def ...uint8) uint8 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Uint8(this.Value)
}

func (this *Var) Uint16(def ...uint16) uint16 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Uint16(this.Value)
}

func (this *Var) Uint32(def ...uint32) uint32 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Uint32(this.Value)
}

func (this *Var) Uint64(def ...uint64) uint64 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Uint64(this.Value)
}

func (this *Var) Float32(def ...float32) float32 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Float32(this.Value)
}

func (this *Var) Float64(def ...float64) float64 {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Float64(this.Value)
}

func (this *Var) Bool(def ...bool) bool {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return Bool(this.Value)
}

func (this *Var) Duration(def ...int) time.Duration {
	if this.IsNil() && len(def) > 0 {
		return time.Duration(def[0])
	}
	return time.Duration(this.Int64())
}

func (this *Var) Microsecond(def ...int) time.Duration {
	return this.Duration(def...) * time.Microsecond
}

func (this *Var) Millisecond(def ...int) time.Duration {
	return this.Duration(def...) * time.Millisecond
}

func (this *Var) Second(def ...int) time.Duration {
	return this.Duration(def...) * time.Second
}

func (this *Var) Minute(def ...int) time.Duration {
	return this.Duration(def...) * time.Minute
}

func (this *Var) Hour(def ...int) time.Duration {
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

func (this *Var) Map(def ...map[string]interface{}) map[string]interface{} {
	return this.GMap(def...)
}

func (this *Var) GMap(def ...map[string]interface{}) map[string]interface{} {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return GMap(this.Value)
}

func (this *Var) SMap(def ...map[string]string) map[string]string {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return SMap(this.Value)
}

func (this *Var) IMap(def ...map[interface{}]interface{}) map[interface{}]interface{} {
	if this.IsNil() && len(def) > 0 {
		return def[0]
	}
	return IMap(this.Value)
}

func (this *Var) DMap(def ...interface{}) *Map {
	if this.IsNil() && len(def) > 0 {
		return NewMap(def[0])
	}
	return NewMap(this.Value)
}
