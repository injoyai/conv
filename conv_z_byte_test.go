package conv

import (
	"testing"
)

func TestInt64ToBytes(t *testing.T) {
	t.Log(Bytes(-100))
	t.Log(Bytes(1000))
	for i := 0; i < 10; i++ {
		t.Log(Bytes(1000))
	}
}

func TestBytesToInt(t *testing.T) {
	bs := Bytes(10000)
	t.Log(bs)
	t.Log(Int(bs))
	t.Log(Int([]byte{255, 254}))
	for i := -100; i < 100; i++ {
		bs2 := Bytes(i)
		num := Int(bs2)
		if num != i {
			t.Error("错误:")
			t.Log(bs2)
			t.Log(num)
		}
	}
}

//0000000000000000
//00000001
func TestBinStr(t *testing.T) {
	t.Log(String(falseArray))
	t.Log(BINStr(65536))
	t.Log(BINStr(256))
	for i := 0; i < 10; i++ {
		t.Log(BINStr(i))
	}
}

func TestInt64ToBinStr(t *testing.T) {
	t.Log(Uint64(-1))
	t.Log(Uint64(-2))
	t.Log(uint64(18446744073709551616 / 2))
	t.Log(BINStr(20))
	t.Log(BINStr(1000))
	t.Log(BINStr(-2))
	t.Log(BINStr(-1))
	t.Log(BINStr(-256))
	t.Log(BINStr(-257))
	t.Log(BINStr(-9223372036854775807))
	t.Log(BINStr(-9223372036854775808))
	//t.Log(BINStr(-9223372036854775809)) //invalid
}

func TestBytesHEX(t *testing.T) {
	t.Log(BytesHEX(20))
	t.Log(BytesHEX(1000))
}

func TestOCT(t *testing.T) {
	t.Log(OCT(20))
	t.Log(OCT(1000)) //1750
	t.Log(OCT(9223372036854775807 - 1))
	t.Log(OCT(-2))
	t.Log(OCT(-1))
}
