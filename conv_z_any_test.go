package conv

import (
	"encoding/hex"
	"log"
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

func TestFloat64bits(t *testing.T) {
	t.Log(Int(Bytes(1.01)))
	t.Log(Uint64(1.01))
	if Int(Bytes(1.01)) != 4607227454796291113 {
		t.Error("错误")
	}
}

func TestCopy(t *testing.T) {

	{ //基础数据的复制
		a := 10
		b := Copy(a)
		if a != b.(int) {
			t.Error("基础数据的复制,测试未通过")
		}
	}
	{ //带指针的基础数据复制
		x := 10
		a := &x
		b := Copy(a)
		t.Logf("%p %p", a, b)
		t.Log(*a, *(b.(*int)))
		if *a != *(b.(*int)) {
			t.Error("带指针的基础数据复制,测试未通过")
		}
	}
	{ //复制对象
		type A struct {
			s string
		}
		a := A{"hello"}
		b := Copy(a).(A)
		b.s = "world"
		t.Log("a:", a)
		t.Log("b:", b)
		if a.s == b.s {
			t.Error("复制对象,测试未通过")
		}
	}
	{ //复制带指针的对象
		type A struct {
			s string
		}
		a := &A{"hello"}
		b := Copy(a).(*A)
		b.s = "world"
		t.Logf("a:%v %p", a, a)
		t.Logf("b:%v %p", b, b)
		if a.s == b.s {
			t.Error("测试未通过")
		}
	}
	{ //复制map
		m := map[interface{}]interface{}{
			"a": 1, 12: 2,
		}
		m2 := Copy(m).(map[interface{}]interface{})
		m2[true] = "hello"
		t.Log(m)
		t.Log(m2)
		if len(m) == len(m2) {
			t.Error("不带指针的Map,测试未通过")
		}
		//带指针的map
		m3 := *(Copy(&m).(*map[interface{}]interface{}))
		m3[true] = "hello"
		t.Log(m)
		t.Log(m3)
		if len(m) == len(m3) {
			t.Error("带指针的Map,测试未通过")
		}
	}
	{ //接口
		a := testI(&testA{})
		a.Set("a")
		b := Copy(a).(testI)
		b.Set("b")
		a.Print()
		b.Print()
		t.Logf("%p,%p", a, b)
		if a == b {
			t.Error("接口复制,测试未通过")
		}
	}
	{ //复制nil
		a := (*testA)(nil)
		b := Copy(a).(*testA)
		t.Logf("%p,%p", a, b)
		t.Log(a == b)
	}

}

type testI interface {
	Set(s string)
	Print()
}

type testA struct {
	s string
}

func (this *testA) Set(s string) {
	this.s = s
}

func (this *testA) Print() {
	log.Println(this.s)
}
