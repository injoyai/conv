package conv

import (
	"math"
)

// toBIN 二进制,返回8倍数大于0长度bool数组,true代表1 根据类型确定长度
func toBIN(i any) (result []bool) {
	bit := IntSize
	switch value := i.(type) {
	case int8, uint8:
		bit = 8
	case int16, uint16:
		bit = 16
	case int32, uint32, float32:
		bit = 32
	case int64, uint64, float64:
		bit = 64
	case int, uint:
	case []byte:
		for _, v := range value {
			result = append(result, toBIN(v)...)
		}
		return
	}
	number := Uint64(i)
	result = make([]bool, bit)
	for x, _ := range result {
		y := uint64(math.Pow(2, float64(bit-x-1)))
		if number >= y {
			result[x] = true
			number -= y
		}
	}
	return result
}

// toOCT int64转8进制,长度22
func toOCT(i any) (result string) {
	number := Uint64(i)
	oct := uint64(8)
	for ; number > 0; number /= oct {
		lsb := octMap[number%oct]
		result = lsb + result
	}
	var s string
	for i := 0; i < 22-len(result); i++ {
		s += "0"
	}
	return s + result
}
