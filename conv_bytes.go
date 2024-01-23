package conv

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"io"
	"math"
	"strconv"
)

type Bs []byte

func (this Bs) Var() *Var {
	return New(this.Bytes())
}

// Len 数据长度
func (this Bs) Len() int {
	return len(this)
}

// Cap cap()
func (this Bs) Cap() int {
	return cap(this)
}

// Error 实现error
func (this Bs) Error() string {
	return this.String()
}

func (this Bs) Copy() Bs {
	cp := Bs(make([]byte, this.Len()))
	copy(cp, this)
	return cp
}

// Append just append
func (this Bs) Append(b ...byte) Bs {
	return append(this, b...)
}

// Upper ASCII小写字母转大小字母
func (this Bs) Upper() Bs {
	return bytes.ToUpper(this)
}

// Lower ASCII大写字母转小写字母
func (this Bs) Lower() Bs {
	return bytes.ToLower(this)
}

// WriteTo 实现io.WriterTo
func (this Bs) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(this)
	return int64(n), err
}

// Equal 判断2个字节数组是否一致
func (this Bs) Equal(bs []byte) bool {
	if (this == nil) != (bs == nil) {
		return false
	}
	return bytes.Equal(this, bs)
}

// Sum 校验和
func (this Bs) Sum() byte {
	b := byte(0)
	for _, v := range this {
		b += v
	}
	return b
}

// Bytes 字节数组
func (this Bs) Bytes() []byte {
	return this
}

// String []{0x31,0x32} >>> "12"
func (this Bs) String() string {
	return string(this)
}

// UTF8 []{0x31,0x32} >>> "12"
func (this Bs) UTF8() string {
	return string(this)
}

// BIN 字节转2进制字符串
func (this Bs) BIN() string {
	return BIN(this)
}

func (this Bs) OCT() string {
	return OCT(this)
}

// HEX []{0x01,0x02} >>> "0102"
func (this Bs) HEX() string {
	return hex.EncodeToString(this)
}

// Base64 same base64.StdEncoding.EncodeToString
func (this Bs) Base64() string {
	return base64.StdEncoding.EncodeToString(this)
}

// HEXBase64 HEX() then Base64()
func (this Bs) HEXBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(this.HEX()))
}

// Reader io.Reader
func (this Bs) Reader() io.Reader {
	return bytes.NewReader(this.Bytes())
}

// Buffer bytes.Buffer
func (this Bs) Buffer() *bytes.Buffer {
	return bytes.NewBuffer(this.Bytes())
}

// GetFirst 获取第一个元素,不存在返回0
func (this Bs) GetFirst() byte {
	if this.Len() > 0 {
		return this[0]
	}
	return 0
}

// GetLast 获取最后一个元素,不存在则返回0
func (this Bs) GetLast() byte {
	if this.Len() > 0 {
		return this[this.Len()-1]
	}
	return 0
}

func (this Bs) Int8() int8 {
	return Int8(this.Bytes())
}

func (this Bs) Int16() int16 {
	return Int16(this.Bytes())
}

func (this Bs) Int() int {
	return Int(this.Bytes())
}

func (this Bs) Int32() int32 {
	return Int32(this.Bytes())
}

// Int64 字节数组转int64 大端模式
func (this Bs) Int64() int64 {
	return Int64(this.Bytes())
}

func (this Bs) Uint8() uint8 {
	return Uint8(this.Bytes())
}

func (this Bs) Uint16() uint16 {
	return Uint16(this.Bytes())
}

func (this Bs) Uint32() uint32 {
	return Uint32(this.Bytes())
}

func (this Bs) Uint() uint {
	return Uint(this.Bytes())
}

// Uint64 字节数组转uint64 大端模式
func (this Bs) Uint64() uint64 {
	return Uint64(this.Bytes())
}

func (this Bs) Float32() float32 {
	return Float32(this.Bytes())
}

func (this Bs) Float64() float64 {
	return Float64(this.Bytes())
}

func (this Bs) Float() float64 {
	return Float64(this.Bytes())
}

func (this Bs) Bool() bool {
	return Bool(this.Bytes())
}

func (this Bs) Float64frombits() float64 {
	return math.Float64frombits(this.Uint64())
}

func (this Bs) Float32frombits() float32 {
	return math.Float32frombits(this.Uint32())
}

// Reverse 倒序
func (this Bs) Reverse() Bs {
	for i := 0; i < this.Len()/2; i++ {
		this[i], this[this.Len()-i-1] = this[this.Len()-i-1], this[i]
	}
	return this
}

// Add 每个字节加add
func (this Bs) Add(add byte) Bs {
	for i, v := range this {
		this[i] = v + add
	}
	return this
}

// Sub 每个字节减sub
func (this Bs) Sub(sub byte) Bs {
	for i, v := range this {
		this[i] = v - sub
	}
	return this
}

// UTF8ToInt []{0x31,0x32} >>> 12
func (this Bs) UTF8ToInt() (int, error) {
	return strconv.Atoi(this.UTF8())
}

// UTF8ToFloat64 字节utf8编码再转int,再转float64
func (this Bs) UTF8ToFloat64(decimals int) (float64, error) {
	i, err := this.UTF8ToInt()
	return float64(i) / math.Pow10(decimals), err
}

// HEXToInt []{0x01,0x02} >>> 102
func (this Bs) HEXToInt() (int, error) {
	return strconv.Atoi(this.HEX())
}

// HEXToFloat64 字节hex编码再转int,再转float64
func (this Bs) HEXToFloat64(decimals int) (float64, error) {
	i, err := this.HEXToInt()
	return float64(i) / math.Pow10(decimals), err
}

// Sub0x33ReverseHEXToInt DLT645协议流程,先减0x33,再倒序,再转hex,再转int
func (this Bs) Sub0x33ReverseHEXToInt() (int, error) {
	return this.Sub(0x33).Reverse().HEXToInt()
}

// Sub0x33ReverseHEXToFloat DLT645协议流程,先减0x33,再倒序,再转hex,再转float64
func (this Bs) Sub0x33ReverseHEXToFloat(decimals int) (float64, error) {
	return this.Sub(0x33).Reverse().HEXToFloat64(decimals)
}
