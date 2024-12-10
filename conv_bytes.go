package conv

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"io"
	"math"
	"strconv"
	"unsafe"
)

type Bytes2 []byte

func (this Bytes2) Var() *Var {
	return New(this.Bytes())
}

// Len 数据长度
func (this Bytes2) Len() int {
	return len(this)
}

// Cap cap()
func (this Bytes2) Cap() int {
	return cap(this)
}

// Error 实现error
func (this Bytes2) Error() string {
	return this.String()
}

// Copy 复制
func (this Bytes2) Copy() Bytes2 {
	cp := Bytes2(make([]byte, this.Len()))
	copy(cp, this)
	return cp
}

// Append just append
func (this Bytes2) Append(b ...byte) Bytes2 {
	return append(this, b...)
}

// Upper ASCII小写字母转大小字母
func (this Bytes2) Upper() Bytes2 {
	return bytes.ToUpper(this)
}

// Lower ASCII大写字母转小写字母
func (this Bytes2) Lower() Bytes2 {
	return bytes.ToLower(this)
}

// WriteTo 实现io.WriterTo
func (this Bytes2) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(this)
	return int64(n), err
}

// Equal 判断2个字节数组是否一致
func (this Bytes2) Equal(bs []byte) bool {
	if (this == nil) != (bs == nil) {
		return false
	}
	return bytes.Equal(this, bs)
}

// Sum 校验和
func (this Bytes2) Sum() byte {
	b := byte(0)
	for _, v := range this {
		b += v
	}
	return b
}

// Bytes 字节数组
func (this Bytes2) Bytes() []byte {
	return this
}

// String []{0x31,0x32} >>> "12"
func (this Bytes2) String() string {
	return string(this)
}

// StringP 获取字符指针(同字节指针)
func (this Bytes2) StringP() *string {
	return (*string)(unsafe.Pointer(&this))
}

// UTF8 []{0x31,0x32} >>> "12"
func (this Bytes2) UTF8() string {
	return string(this)
}

// BIN 字节转2进制字符串
func (this Bytes2) BIN() string {
	return BIN(this)
}

func (this Bytes2) OCT() string {
	return OCT(this)
}

// HEX []{0x01,0x02} >>> "0102"
func (this Bytes2) HEX() string {
	return hex.EncodeToString(this)
}

// Base64 same base64.StdEncoding.EncodeToString
func (this Bytes2) Base64() string {
	return base64.StdEncoding.EncodeToString(this)
}

// HEXBase64 HEX() then Base64()
func (this Bytes2) HEXBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(this.HEX()))
}

// Reader io.Reader
func (this Bytes2) Reader() io.Reader {
	return bytes.NewReader(this.Bytes())
}

// Buffer bytes.Buffer
func (this Bytes2) Buffer() *bytes.Buffer {
	return bytes.NewBuffer(this.Bytes())
}

// GetFirst 获取第一个元素,不存在返回0
func (this Bytes2) GetFirst() byte {
	if this.Len() > 0 {
		return this[0]
	}
	return 0
}

// GetLast 获取最后一个元素,不存在则返回0
func (this Bytes2) GetLast() byte {
	if this.Len() > 0 {
		return this[this.Len()-1]
	}
	return 0
}

func (this Bytes2) Int8() int8 {
	return Int8(this.Bytes())
}

func (this Bytes2) Int16() int16 {
	return Int16(this.Bytes())
}

func (this Bytes2) Int() int {
	return Int(this.Bytes())
}

func (this Bytes2) Int32() int32 {
	return Int32(this.Bytes())
}

// Int64 字节数组转int64 大端模式
func (this Bytes2) Int64() int64 {
	return Int64(this.Bytes())
}

func (this Bytes2) Uint8() uint8 {
	return Uint8(this.Bytes())
}

func (this Bytes2) Uint16() uint16 {
	return Uint16(this.Bytes())
}

func (this Bytes2) Uint32() uint32 {
	return Uint32(this.Bytes())
}

func (this Bytes2) Uint() uint {
	return Uint(this.Bytes())
}

// Uint64 字节数组转uint64 大端模式
func (this Bytes2) Uint64() uint64 {
	return Uint64(this.Bytes())
}

func (this Bytes2) Float32() float32 {
	return Float32(this.Bytes())
}

func (this Bytes2) Float64() float64 {
	return Float64(this.Bytes())
}

func (this Bytes2) Float() float64 {
	return Float64(this.Bytes())
}

func (this Bytes2) Bool() bool {
	return Bool(this.Bytes())
}

func (this Bytes2) Float64frombits() float64 {
	return math.Float64frombits(this.Uint64())
}

func (this Bytes2) Float32frombits() float32 {
	return math.Float32frombits(this.Uint32())
}

// Reverse 倒序
func (this Bytes2) Reverse() Bytes2 {
	for i := 0; i < this.Len()/2; i++ {
		this[i], this[this.Len()-i-1] = this[this.Len()-i-1], this[i]
	}
	return this
}

// Add 每个字节加add
func (this Bytes2) Add(add byte) Bytes2 {
	for i, v := range this {
		this[i] = v + add
	}
	return this
}

// Sub 每个字节减sub
func (this Bytes2) Sub(sub byte) Bytes2 {
	for i, v := range this {
		this[i] = v - sub
	}
	return this
}

// UTF8ToInt []{0x31,0x32} >>> 12
func (this Bytes2) UTF8ToInt() (int, error) {
	return strconv.Atoi(this.UTF8())
}

// UTF8ToFloat64 字节utf8编码再转int,再转float64
func (this Bytes2) UTF8ToFloat64(decimals int) (float64, error) {
	i, err := this.UTF8ToInt()
	return float64(i) / math.Pow10(decimals), err
}

// HEXToInt []{0x01,0x02} >>> 102
func (this Bytes2) HEXToInt() (int, error) {
	return strconv.Atoi(this.HEX())
}

// HEXToFloat64 字节hex编码再转int,再转float64
func (this Bytes2) HEXToFloat64(decimals int) (float64, error) {
	i, err := this.HEXToInt()
	return float64(i) / math.Pow10(decimals), err
}

// Sub0x33ReverseHEXToInt DLT645协议流程,先减0x33,再倒序,再转hex,再转int
func (this Bytes2) Sub0x33ReverseHEXToInt() (int, error) {
	return this.Sub(0x33).Reverse().HEXToInt()
}

// Sub0x33ReverseHEXToFloat DLT645协议流程,先减0x33,再倒序,再转hex,再转float64
func (this Bytes2) Sub0x33ReverseHEXToFloat(decimals int) (float64, error) {
	return this.Sub(0x33).Reverse().HEXToFloat64(decimals)
}
