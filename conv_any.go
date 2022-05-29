package conv

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
func Uint32(i interface{}) uint32 {
	return uint32(Uint64(i))
}

// Uint64 任意类型转 uint64.
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

// Interfaces 任意类型转 []Interface.
func Interfaces(i interface{}) []interface{} {
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

// OCT 任意类型转 string 长度固定22,8进制,'0'-'7'.
// -1 >>> "1777777777777777777777"
// +1 >>> "0000000000000000000001"
func OCT(i interface{}) string {
	return toOCT(i)
}
