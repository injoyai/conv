package conv

import (
	"encoding/hex"
	"reflect"
)

// Byte 任意类型转 byte.
func Byte(i interface{}) byte {
	return Uint8(i)
}

// Bytes 任意类型转 []byte.
func Bytes(i interface{}) Bs {
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
// 去除支持IEEE754标准 1.01 >>> 1094057984 使用 math.Float32bits(value)
func Uint32(i interface{}) uint32 {
	return uint32(Uint64(i))
}

// Uint64 任意类型转 uint64.
// 去除支持IEEE754标准 1.01 >>> 4622593173774925824 使用 math.Float64bits(value)
func Uint64(i interface{}) uint64 {
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

// SMap 任意类型转 map[string]string
func SMap(i interface{}) map[string]string {
	return toSMap(i)
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

// Copy 复制任意数据
func Copy(i interface{}) interface{} {
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
		copyValue(result, original)
		return result.Interface()
	}
}

// copyValue 递归复制值，处理指针的情况
func copyValue(result, original reflect.Value) {
	switch original.Kind() {
	case reflect.Ptr:
		// 如果是指针，则递归复制指针指向的内容
		// result的值暂时还是nil,无需判断
		if !original.IsNil() {
			result.Set(reflect.New(original.Type().Elem()))
			copyValue(result.Elem(), original.Elem())
		}
	case reflect.Struct:
		// 如果是结构体，则递归复制结构体的字段
		for i := 0; i < original.NumField(); i++ {
			if result.Field(i).CanSet() {
				copyValue(result.Field(i), original.Field(i))
			}
		}
	case reflect.Map:
		if !original.IsNil() {
			// 如果是映射，则创建新的映射并递归复制键值对
			result.Set(reflect.MakeMap(original.Type()))
			for _, key := range original.MapKeys() {
				destKey := reflect.New(key.Type()).Elem()
				copyValue(destKey, key)
				destValue := reflect.New(original.MapIndex(key).Type()).Elem()
				copyValue(destValue, original.MapIndex(key))
				result.SetMapIndex(destKey, destValue)
			}
		}
	case reflect.Slice:
		if !original.IsNil() {
			// 如果是切片，则创建新的切片并递归复制元素
			result.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
			for i := 0; i < original.Len(); i++ {
				copyValue(result.Index(i), original.Index(i))
			}
		}
	default:
		result.Set(original)
	}
}
