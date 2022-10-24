package conv

import (
	"encoding/hex"
	"math"
	"testing"
	"time"
)

func TestBytes(t *testing.T) {
	t.Log(hex.EncodeToString(Bytes(100))) // 64
	t.Log(Byte("00000100"))               // 64
	t.Log(Bytes(int(100)))
}

func TestInt64(t *testing.T) {
	t.Log(New("1h3m20s").Duration())
	t.Log(New("1h3m20s").Duration() / time.Second)
	t.Log(Int64("1h3m20s"))
	t.Log(Int64("0x20"))  // 32
	t.Log(Int64("0X20"))  // 32
	t.Log(Int64("0b20"))  // 0
	t.Log(Int64("0b100")) // 4
	t.Log(Int64("0B100")) // 4
	t.Log(Int64("020"))   // 16
	t.Log(Int64("20"))    //20

	t.Log(Uint64("0x20"))  // 32
	t.Log(Uint64("0X20"))  // 32
	t.Log(Uint64("0b20"))  // 0
	t.Log(Uint64("0b100")) // 4
	t.Log(Uint64("0B100")) // 4
	t.Log(Uint64("020"))   // 16
	t.Log(Uint64("20"))    //20

	t.Log(uint64(1 << 8))
	t.Log(uint64(1 << 63))
	t.Log(Uint64(-1))
	t.Log(uint64(math.MaxUint64))

	t.Log(Bytes(uint64(math.MaxUint64)))
	t.Log(Byte(uint64(math.MaxUint64 - 1)))
	t.Log(Byte("010"))       // 8 八进制
	t.Log(Byte("0b10"))      // 2 二进制
	t.Log(Byte("0x10"))      // 16 十六进制
	t.Log(Byte("10"))        // 10 十进制
	t.Log(Bytes("0000010"))  // [48 48 48 48 48 49 48]
	t.Log(Bytes("00000010")) // [48 48 48 48 48 48 49 48]

}

func TestUint(t *testing.T) {
	bs := Bytes(float32(11.375))
	t.Log(hex.EncodeToString(bs))
	t.Log(Int(bs))

	if Uint32(11.375) != 1094057984 {
		t.Log(Bytes(11.375))
		t.Log(Uint64(Bytes(11.375)))
		t.Log(Uint32(Bytes(11.375)))
		t.Error("错误")
		return
	}
	if Uint64(11.375) != 4622593173774925824 {
		t.Error("错误")
		return
	}
	if Uint32(float32(11.375)) != 1094057984 {
		t.Error("错误")
		return
	}
	if Uint64(float32(11.375)) != 4622593173774925824 {
		t.Error("错误")
		return
	}
	if Float32(Bytes(uint64(4622593173774925824))) != 11.375 {
		t.Error("错误")
		return
	}
	if Float32(uint32(1094057984)) != 11.375 {
		t.Log(Float32(uint32(1094057984)))
		t.Error("错误")
		return
	}
	if Float64(uint32(1094057984)) != 11.375 {
		t.Error("错误")
		return
	}
	if Float64(uint64(4622593173774925824)) != 11.375 {
		t.Error("错误")
		return
	}
}

func TestMap(t *testing.T) {

}
