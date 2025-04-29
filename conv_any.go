package conv

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"time"
)

// Byte 任意类型转 byte.
func Byte(i any) byte {
	return Uint8(i)
}

// Bytes 任意类型转 []byte.
func Bytes(i any) []byte {
	return toBytes(i)
}

// BytesZ 字节拓展
func BytesZ(i any) Bytes2 {
	return toBytes(i)
}

// Rune 任意类型转 rune.
func Rune(i any) rune {
	return Int32(i)
}

// Runes 任意类型转 []rune.
func Runes(i any) []rune {
	return []rune(String(i))
}

// String 任意类型转 string.
func String(i any) string {
	return toString(i)
}

// Strings 任意类型转 []string
func Strings(i any) (list []string) {
	for _, v := range Interfaces(i) {
		list = append(list, String(v))
	}
	return
}

// Int 任意类型转 int.
func Int(i any) int {
	return int(Int64(i))
}

// Ints 任意类型转 []int.
func Ints(i any) (list []int) {
	for _, v := range Interfaces(i) {
		list = append(list, Int(v))
	}
	return
}

// Int8 任意类型转 int8.
func Int8(i any) int8 {
	return int8(Int64(i))
}

// Int16 任意类型转 int16.
func Int16(i any) int16 {
	return int16(Int64(i))
}

// Int32 任意类型转 int32.
func Int32(i any) int32 {
	return int32(Int64(i))
}

// Int64 任意类型转 int64.
func Int64(i any) int64 {
	return toInt64(i)
}

// Int64s 任意类型转 []int64.
func Int64s(i any) (list []int64) {
	for _, v := range Interfaces(i) {
		list = append(list, Int64(v))
	}
	return
}

// Uint 任意类型转 uint.
func Uint(i any) uint {
	return uint(Uint64(i))
}

// Uint8 任意类型转 uint8.
func Uint8(i any) uint8 {
	return uint8(Uint64(i))
}

// Uint16 任意类型转 uint16.
func Uint16(i any) uint16 {
	return uint16(Uint64(i))
}

// Uint32 任意类型转 uint32.
// 去除支持IEEE754标准 1.01 >>> 1094057984 使用 math.Float32bits(value)
func Uint32(i any) uint32 {
	return uint32(Uint64(i))
}

// Uint64 任意类型转 uint64.
// 去除支持IEEE754标准 1.01 >>> 4622593173774925824 使用 math.Float64bits(value)
func Uint64(i any) uint64 {
	return toUint64(i)
}

// Float32 任意类型转 float32.
func Float32(i any) float32 {
	return float32(Float64(i))
}

// Float64 任意类型转 float64.
func Float64(i any) float64 {
	return toFloat64(i)
}

// Bool 任意类型转 bool.
func Bool(i any) bool {
	return toBool(i)
}

// GMap 任意类型转 map[string]any
func GMap(i any) map[string]any {
	return toGMap(i)
}

// SMap 任意类型转 map[string]string
func SMap(i any) map[string]string {
	return toSMap(i)
}

// IMap 任意类型转 map[any]any
func IMap(i any) map[any]any {
	return toIMap(i)
}

// DMap 解析任意数据
func DMap(i any) *Map {
	return NewMap(i)
}

// Interfaces 任意类型转 []Interface.
func Interfaces(i any) []any {
	return toInterfaces(i)
}

// Array 任意类型转 []Interface.
func Array(i any) []any {
	return toInterfaces(i)
}

// Duration 任意类型转 time.Duration
func Duration(i any) time.Duration {
	return time.Duration(Int64(i))
}

// BINBool 任意类型转二进制 []bool 返回长度8的倍数且大于0,true代表二进制的1.
func BINBool(i any) []bool {
	return toBIN(i)
}

// BIN 任意类型转 []bool 返回长度8的倍数且大于0,true代表二进制的1.
// +1 >>> "00000001"
// -1 >>> "11111111"
func BIN(i any) string {
	result := ""
	for _, v := range toBIN(i) {
		result += func() string {
			if v {
				return "1"
			}
			return "0"
		}()
	}
	return result
}

// BINStr 任意类型转 string 长度8的倍数且大于0,由'1'和'0'组成,
// +1 >>> "00000001"
// -1 >>> "11111111"
func BINStr(i any) string {
	return BIN(i)
}

// OCT 任意类型转8进制 string 长度固定22,8进制,'0'-'7'.
// -1 >>> "1777777777777777777777"
// +1 >>> "0000000000000000000001"
func OCT(i any) string {
	return toOCT(i)
}

// OCTStr 任意类型转 string 长度固定22,8进制,'0'-'7'.
// -1 >>> "1777777777777777777777"
// +1 >>> "0000000000000000000001"
func OCTStr(i any) string {
	return toOCT(i)
}

// HEX 转16进制字符串
func HEX(i any) string {
	return hex.EncodeToString(Bytes(i))
}

// HEXStr 转16进制字符串
func HEXStr(i any) string {
	return hex.EncodeToString(Bytes(i))
}

// Unmarshal 任意类型i转到ptr
func Unmarshal(i any, ptr any, param ...UnmarshalParam) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	return unmarshal(i, ptr, param...)
}

// Copy 复制任意数据
func Copy(i any) any {
	if i == nil {
		return nil
	}
	switch v := i.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64, complex64, complex128,
		string, bool:
		return v
	case apiCopy:
		return v.Copy()
	default:
		original := reflect.ValueOf(i)
		result := reflect.New(original.Type()).Elem()
		copySameKind(result, original)
		return result.Interface()
	}
}
