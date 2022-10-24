package conv

import (
	"encoding/hex"
	"math"
)

// Byte 任意类型转 byte.
func Byte(i interface{}) byte {
	return Uint8(i)
}

// Bytes 任意类型转 []byte.
func Bytes(i interface{}) []byte {
	return toBytes(i)
}

// Rune 任意类型转 rune.
func Rune(i interface{}) rune {
	return Int32(i)
}

// Runes 任意类型转 []rune.
func Runes(i interface{}) []rune {
	return []rune(String(i))
}

// String 任意类型转 string.
func String(i interface{}) string {
	return toString(i)
}

// Strings 任意类型转 []string
func Strings(i interface{}) (list []string) {
	for _, v := range Interfaces(i) {
		list = append(list, String(v))
	}
	return
}

// Int 任意类型转 int.
func Int(i interface{}) int {
	return int(Int64(i))
}

// Ints 任意类型转 []int.
func Ints(i interface{}) (list []int) {
	for _, v := range Interfaces(i) {
		list = append(list, Int(v))
	}
	return
}

// Int8 任意类型转 int8.
func Int8(i interface{}) int8 {
	return int8(Int64(i))
}

// Int16 任意类型转 int16.
func Int16(i interface{}) int16 {
	return int16(Int64(i))
}

// Int32 任意类型转 int32.
func Int32(i interface{}) int32 {
	return int32(Int64(i))
}

// Int64 任意类型转 int64.
func Int64(i interface{}) int64 {
	return toInt64(i)
}

// Int64s 任意类型转 []int64.
func Int64s(i interface{}) (list []int64) {
	for _, v := range Interfaces(i) {
		list = append(list, Int64(v))
	}
	return
}

// Uint 任意类型转 uint.
func Uint(i interface{}) uint {
	return uint(Uint64(i))
}

// Uint8 任意类型转 uint8.
func Uint8(i interface{}) uint8 {
	return uint8(Uint64(i))
}

// Uint16 任意类型转 uint16.
func Uint16(i interface{}) uint16 {
	return uint16(Uint64(i))
}

// Uint32 任意类型转 uint32.
// 支持IEEE754标准 1.01 >>> 1094057984
func Uint32(i interface{}) uint32 {
	switch value := i.(type) {
	case float32:
		return math.Float32bits(value)
	case float64:
		return Uint32(float32(value))
	}
	return uint32(Uint64(i))
}

// Uint64 任意类型转 uint64.
// 支持IEEE754标准 1.01 >>> 4622593173774925824
func Uint64(i interface{}) uint64 {
	switch value := i.(type) {
	case float32:
		return Uint64(float64(value))
	case float64:
		return math.Float64bits(value)
	}
	return toUint64(i)
}

// Float32 任意类型转 float32.
func Float32(i interface{}) float32 {
	return float32(Float64(i))
}

// Float64 任意类型转 float64.
func Float64(i interface{}) float64 {
	return toFloat64(i)
}

// Bool 任意类型转 bool.
func Bool(i interface{}) bool {
	return toBool(i)
}

// GMap 任意类型转 map[string]interface{}
func GMap(i interface{}) map[string]interface{} {
	return toGMap(i)
}

// IMap 任意类型转 map[interface{}]interface{}
func IMap(i interface{}) map[interface{}]interface{} {
	return toIMap(i)
}

// DMap 解析任意数据
func DMap(i interface{}) *Map {
	return NewMap(i)
}

// Interfaces 任意类型转 []Interface.
func Interfaces(i interface{}) []interface{} {
	return toInterfaces(i)
}

// Array 任意类型转 []Interface.
func Array(i interface{}) []interface{} {
	return toInterfaces(i)
}

// BIN 任意类型转 []bool 返回长度8的倍数且大于0,true代表二进制的1.
func BIN(i interface{}) []bool {
	return toBIN(i)
}

// BINStr 任意类型转 string 长度8的倍数且大于0,由'1'和'0'组成,
// +1 >>> "00000001"
// -1 >>> "11111111"
func BINStr(i interface{}) string {
	result := ""
	for _, v := range BIN(i) {
		result += func() string {
			if v {
				return "1"
			}
			return "0"
		}()
	}
	return result
}

// OCTStr 任意类型转 string 长度固定22,8进制,'0'-'7'.
// -1 >>> "1777777777777777777777"
// +1 >>> "0000000000000000000001"
func OCTStr(i interface{}) string {
	return toOCT(i)
}

// HEXStr 转16进制字符串
func HEXStr(i interface{}) string {
	return hex.EncodeToString(Bytes(i))
}

// HEXBytes 16进制的方式转字节
func HEXBytes(i interface{}) []byte {
	bs, _ := hex.DecodeString(String(i))
	return bs
}

// HEXInt 先16进制转字节,在转Int
func HEXInt(i interface{}) int {
	return Int(HEXBytes(i))
}
