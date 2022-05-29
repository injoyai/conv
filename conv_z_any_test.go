package conv

import (
	"encoding/hex"
	"math"
	"testing"
)

func TestBytes(t *testing.T) {
	t.Log(hex.EncodeToString(Bytes(100)))
	t.Log(Byte("00000100"))
	x := falseArray
	x[4] = true
	t.Log(Byte(x))
}

func TestInt64(t *testing.T) {
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
	t.Log(Byte("0000010"))    // 8 八进制
	t.Log(Byte("0b00000010")) // 2 二进制
	t.Log(Bytes("0000010"))   // [48 48 48 48 48 49 48]
	t.Log(Bytes("00000010"))

}
