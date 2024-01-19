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

func (this Bs) Cap() int {
	return cap(this)
}

func (this Bs) Error() string {
	return this.String()
}

func (this Bs) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(this)
	return int64(n), err
}

func (this Bs) Sum() byte {
	b := byte(0)
	for _, v := range this {
		b += v
	}
	return b
}

func (this Bs) Copy() Bs {
	cp := make([]byte, len(this))
	copy(cp, this)
	return cp
}

func (this Bs) Upper() Bs {
	return bytes.ToUpper(this)
}

func (this Bs) Lower() Bs {
	return bytes.ToLower(this)
}

// String []{0x31,0x32} >>> "12"
func (this Bs) String() string {
	return string(this)
}

// UTF8 []{0x31,0x32} >>> "12"
func (this Bs) UTF8() string {
	return string(this)
}

// ASCII []{0x31,0x32} >>> "12"
func (this Bs) ASCII() string {
	return string(this)
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
	return Bs(this.HEX()).Base64()
}

// Bytes 字节数组
func (this Bs) Bytes() []byte {
	return this
}

// Reader io.Reader
func (this Bs) Reader() io.Reader {
	return bytes.NewReader(this.Bytes())
}

// Buffer bytes.Buffer
func (this Bs) Buffer() *bytes.Buffer {
	return bytes.NewBuffer(this.Bytes())
}

// SumByte 累加转byte
func (this Bs) SumByte() byte {
	b := byte(0)
	for _, v := range this {
		b += v
	}
	return b
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

// BINStr 字节转2进制字符串
func (this Bs) BINStr() string {
	return BINStr(this)
}

// BIN 字节转2进制字符串
func (this Bs) BIN() string {
	return BINStr(this)
}

// Append just append
func (this Bs) Append(b ...byte) Bs {
	return append(this, b...)
}

// UTF8ToInt []{0x31,0x32} >>> 12
func (this Bs) UTF8ToInt() (int, error) {
	return strconv.Atoi(this.ASCII())
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

// Reverse 倒序
func (this Bs) Reverse() Bs {
	x := make([]byte, len(this))
	for i, v := range this {
		x[len(this)-i-1] = v
	}
	return x
}

// ReverseASCII 倒序再ASCII
func (this Bs) ReverseASCII() string {
	return this.Reverse().ASCII()
}

// ReverseHEX 倒序再hex
func (this Bs) ReverseHEX() string {
	return this.Reverse().HEX()
}

// ReverseBase64 倒序再base64
func (this Bs) ReverseBase64() string {
	return this.Reverse().Base64()
}

// AddByte 每个字节加add
func (this Bs) AddByte(add byte) Bs {
	result := make([]byte, len(this))
	for _, v := range this {
		result = append(result, v+add)
	}
	return result
}

// SubByte 每个字节减sub
func (this Bs) SubByte(sub byte) Bs {
	result := make([]byte, len(this))
	for _, v := range this {
		result = append(result, v-sub)
	}
	return result
}

// Sub0x33 每个字节减0x33
func (this Bs) Sub0x33() Bs {
	return this.SubByte(0x33)
}

// Add0x33 每个字节加0x33
func (this Bs) Add0x33() Bs {
	return this.AddByte(0x33)
}

// Sub0x33ReverseHEXToInt DLT645协议流程,先减0x33,再倒序,再转hex,再转int
func (this Bs) Sub0x33ReverseHEXToInt() (int, error) {
	return this.Sub0x33().Reverse().HEXToInt()
}

// Sub0x33ReverseHEXToFloat DLT645协议流程,先减0x33,再倒序,再转hex,再转float64
func (this Bs) Sub0x33ReverseHEXToFloat(decimals int) (float64, error) {
	return this.Sub0x33().Reverse().HEXToFloat64(decimals)
}
