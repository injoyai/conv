package conv

import (
	"encoding/hex"
	"math"
)

// toBIN 二进制,返回8倍数大于0长度bool数组,true代表1
func toBIN(i interface{}) (result []bool) {
	number := Uint64(i)
	bit := 64
	if number > uint64(1<<63) {
		x := math.MaxUint64 - number + 1
		bit = 8
		for x >= 1<<8 {
			x /= 1 << 8
			bit += 8
		}
		switch bit {
		case 8:
			number = Uint64(Uint8(i))
		case 16:
			number = Uint64(Uint16(i))
		case 24, 32:
			number = Uint64(Uint32(i))
		case 40, 48, 56, 64:
		}
	}
	result = make([]bool, bit)
	for i, _ := range result {
		x := uint64(math.Pow(2, float64(bit-i-1)))
		if number >= x {
			result[i] = true
			number -= x
		}
	}
	for len(result) > 8 && String(result[:8]) == String(falseArray) {
		result = result[8:]
	}
	return
}

// toOCT int64转8进制,长度22
func toOCT(i interface{}) (result string) {
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

// BytesHEX 转16进制字符串
func BytesHEX(i interface{}) string {
	return hex.EncodeToString(Bytes(i))
}
