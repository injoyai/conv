package conv

import (
	"encoding/hex"
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
}
