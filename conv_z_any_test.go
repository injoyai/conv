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
	{ //空指针
		m := map[interface{}]interface{}(nil)
		m2 := Copy(m).(map[interface{}]interface{})
		t.Log(m == nil)
		t.Log(m2 == nil)
		s := []interface{}(nil)
		s2 := Copy(s).([]interface{})
		t.Log(s == nil)
		t.Log(s2 == nil)
		a := (*testA)(nil)
		b := Copy(a).(*testA)
		t.Log(a == nil)
		t.Log(b == nil)
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

func Test_unmarshal(t *testing.T) {
	type (
		_any    interface{}
		_struct struct {
			A string
		}
		_map map[string]interface{}
	)

	valList := []interface{}{
		nil, "a", 1, true, 1.02, _any(8), _any(nil), _struct{}, _map{"s": 1.1}, _map(nil),
	}

	{
		t.Log("=============测试String=============")
		for _, v := range valList {
			ptr := "s"
			if err := unmarshal(v, &ptr); err != nil {
				t.Error(err)
				return
			}
			t.Logf("%T(%#v) > %T(%#v)", v, v, ptr, ptr)
			if v != nil && String(ptr) != String(v) {
				t.Error("测试失败,结果不一致")
				return
			}
			t.Log("测试通过---")
		}
		t.Log("=============测试结束=============")
	}

	{
		t.Log("=============测试Int=============")
		for _, v := range valList {
			ptr := 8
			if err := unmarshal(v, &ptr); err != nil {
				t.Error(err)
				return
			}
			t.Logf("%T(%#v) > %T(%#v)", v, v, ptr, ptr)
			if v != nil && Int(ptr) != Int(v) {
				t.Error("测试失败,结果不一致")
				return
			}
			t.Log("测试通过---")
		}
	}

	{
		t.Log("=============测试Slice=============")
		for _, v := range valList {
			ptr := []string{"5"}
			if err := unmarshal(v, &ptr); err != nil {
				t.Error(err)
				return
			}
			t.Logf("%T(%#v) > %T(%#v)", v, v, ptr, ptr)
			if v != nil && String(Strings(ptr)) != String(Strings(v)) {
				t.Error("测试失败,结果不一致", String(Strings(v)), " != ", String(Strings(ptr)))
				return
			}
			t.Log("测试通过---")
		}
	}

	t.Log("=============测试全部结束=============")
}

func Test_MapToStruct(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": "2",
		"c": true,
		"d": 1.02,
		"e": 20.1,
		"f": 30.6,
		"G": map[string]interface{}{
			"h": "10",
		},
	}
	type _struct struct {
		A int     `json:"a"`
		B string  `json:"b"`
		C bool    `json:"c"`
		D float64 `json:"d"`
		E float64
		F string `json:"f"`
		G struct {
			H int `json:"h"`
		}
	}
	type _struct2 struct {
		A int    `json:"a"`
		B string `json:"b"`
	}
	{
		x := &_struct2{A: 100}
		if err := unmarshal(m, x); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v", *x)
	}
	{
		x := new(_struct)
		if err := unmarshal(m, x); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v", *x)
	}
	{
		var x *_struct
		if err := unmarshal(m, &x); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v", x)
	}
	{
		var x **_struct
		if err := unmarshal(m, &x); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v", *x)
	}
}

func Test_StructToMap(t *testing.T) {
	type _struct struct {
		A int     `json:"a"`
		B string  `json:"b"`
		C bool    `json:"c"`
		D float64 `json:"d"`
		E float64
		F string `json:"f"`
		G struct {
			H int `json:"h"`
		}
	}
	x := _struct{
		A: 1,
		B: "2",
		C: true,
		D: 1.02,
		E: 20.1,
		F: "30.6",
		G: struct {
			H int `json:"h"`
		}{
			H: 10,
		},
	}
	{
		m := map[string]interface{}{}
		if err := unmarshal(x, &m); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v", m)
	}
	{
		m := map[string]string{}
		if err := unmarshal(x, &m); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v", m)
	}
	{
		if err := unmarshal(x, nil); err != nil {
			t.Error(err)
			return
		}
	}
}

func Test_MapToMap(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": "2",
		"c": true,
		"d": 1.02,
		"e": 20.1,
		"f": 30.6,
		"G": map[string]interface{}{
			"h": "10",
		},
	}
	{
		m2 := map[string]string(nil)
		if err := unmarshal(m, &m2); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v", m2)
	}
	{
		var m2 map[string]string
		if err := unmarshal(m, &m2); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v", m2)
	}
	{
		var m2 map[interface{}]string
		if err := unmarshal(m, &m2); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v", m2)
	}
	{
		m2 := map[interface{}]string{1: "1"}
		if err := unmarshal(m, &m2); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v", m2)
	}
}
