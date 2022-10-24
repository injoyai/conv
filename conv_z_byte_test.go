package conv

import (
	"testing"
)

func TestInt64ToBytes(t *testing.T) {
	t.Log(Bytes(uint16(0)))
	t.Log(Bytes(uint8(0)))
	t.Log(Bytes(int16(-1)))
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
	t.Log(BINStr(65536))
	t.Log(BINStr(256))
	t.Log(BINStr(-1))
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
	t.Log(2e2)
	t.Log(10e2)
	t.Log(HEXStr(20))
	t.Log(HEXStr(1000))
	t.Log(HEXStr(uint16(1000)))
	x := HEXStr(1999)
	t.Log(x)
	t.Log(HEXInt(x))
	t.Log(Int(x))
}

func TestOCT(t *testing.T) {
	t.Log(OCTStr(20))
	t.Log(OCTStr(1000)) //1750
	t.Log(OCTStr(9223372036854775807 - 1))
	t.Log(OCTStr(-2))
	t.Log(OCTStr(-1))
}

func TestBIN(t *testing.T) {
	t.Log(toBIN(10))
	t.Log(BINStr(10))
	t.Log(BINStr(-1))        //
	t.Log(BINStr(int32(-1))) //
	t.Log(BINStr(int8(-1)))  //11111111
	t.Log(int8(-1))
	t.Log(Uint64(int8(-1)))
}

func TestBIN2(t *testing.T) {
	t.Log(BINStr(-1.1))
}
