package conv

import "time"

type Extend interface {
	IGetVar
	IsDefault(key string) bool
	IsNumber(key string) bool
	IsString(key string) bool
	IsInt(key string) bool
	IsFloat(key string) bool
	IsBool(key string) bool
	IsTime(key string) bool
	IsArray(key string) bool
	IsPointer(key string) bool
	IsNil(key string) bool
	Get(key string) *Var
	GetInterface(key string, def ...interface{}) interface{}
	GetInterfaces(key string, def ...[]interface{}) []interface{}
	GetByte(key string, def ...byte) byte
	GetBytes(key string, def ...[]byte) []byte
	GetRune(key string, def ...rune) rune
	GetRunes(key string, def ...[]rune) []rune
	GetString(key string, def ...string) string
	GetStrings(key string, def ...[]string) []string
	GetInt(key string, def ...int) int
	GetInts(key string, def ...[]int) []int
	GetInt8(key string, def ...int8) int8
	GetInt16(key string, def ...int16) int16
	GetInt32(key string, def ...int32) int32
	GetInt64(key string, def ...int64) int64
	GetInt64s(key string, def ...[]int64) []int64
	GetUint(key string, def ...uint) uint
	GetUint8(key string, def ...uint8) uint8
	GetUint16(key string, def ...uint16) uint16
	GetUint32(key string, def ...uint32) uint32
	GetUint64(key string, def ...uint64) uint64
	GetFloat32(key string, def ...float32) float32
	GetFloat64(key string, def ...float64) float64
	GetBool(key string, def ...bool) bool
	GetDuration(key string, def ...time.Duration) time.Duration
	GetMicrosecond(key string, def ...time.Duration) time.Duration
	GetMillisecond(key string, def ...time.Duration) time.Duration
	GetSecond(key string, def ...time.Duration) time.Duration
	GetMinute(key string, def ...time.Duration) time.Duration
	GetHour(key string, def ...time.Duration) time.Duration
	GetErr(key string, def ...error) error
	GetGMap(key string, def ...map[string]interface{}) map[string]interface{}
	GetSMap(key string, def ...map[string]string) map[string]string
	GetIMap(key string, def ...map[interface{}]interface{}) map[interface{}]interface{}
	GetDMap(key string, def ...interface{}) *Map
}

type IGetVar interface {
	GetVar(string) *Var
}

func NewExtend(i IGetVar) Extend {
	return &extend{IGetVar: i}
}

type extend struct {
	IGetVar
}

func (this *extend) IsDefault(key string) bool {
	return IsDefault(this.GetVar(key).Value)
}

func (this *extend) IsNumber(key string) bool {
	return IsNumber(this.GetVar(key).Value)
}

func (this *extend) IsString(key string) bool {
	return IsString(this.GetVar(key).Value)
}

func (this *extend) IsInt(key string) bool {
	return IsInt(this.GetVar(key).Value)
}

func (this *extend) IsFloat(key string) bool {
	return IsFloat(this.GetVar(key).Value)
}

func (this *extend) IsBool(key string) bool {
	return IsBool(this.GetVar(key).Value)
}

func (this *extend) IsTime(key string) bool {
	return IsTime(this.GetVar(key).Value)
}

func (this *extend) IsArray(key string) bool {
	return IsArray(this.GetVar(key).Value)
}

func (this *extend) IsPointer(key string) bool {
	return IsPointer(this.GetVar(key).Value)
}

func (this *extend) IsNil(key string) bool {
	return IsNil(this.GetVar(key).Value)
}

func (this *extend) Get(key string) *Var {
	return this.GetVar(key)
}

func (this *extend) GetInterface(key string, def ...interface{}) interface{} {
	return this.GetVar(key).Interface(def...)
}

func (this *extend) GetInterfaces(key string, def ...[]interface{}) []interface{} {
	return this.GetVar(key).Interfaces(def...)
}

func (this *extend) GetByte(key string, def ...byte) byte {
	return this.GetVar(key).Byte(def...)
}

func (this *extend) GetBytes(key string, def ...[]byte) []byte {
	return this.GetVar(key).Bytes(def...)
}

func (this *extend) GetRune(key string, def ...rune) rune {
	return this.GetVar(key).Rune(def...)
}

func (this *extend) GetRunes(key string, def ...[]rune) []rune {
	return this.GetVar(key).Runes(def...)
}

func (this *extend) GetString(key string, def ...string) string {
	return this.GetVar(key).String(def...)
}

func (this *extend) GetStrings(key string, def ...[]string) []string {
	return this.GetVar(key).Strings(def...)
}

func (this *extend) GetInt(key string, def ...int) int {
	return this.GetVar(key).Int(def...)
}

func (this *extend) GetInts(key string, def ...[]int) []int {
	return this.GetVar(key).Ints(def...)
}

func (this *extend) GetInt8(key string, def ...int8) int8 {
	return this.GetVar(key).Int8(def...)
}

func (this *extend) GetInt16(key string, def ...int16) int16 {
	return this.GetVar(key).Int16(def...)
}

func (this *extend) GetInt32(key string, def ...int32) int32 {
	return this.GetVar(key).Int32(def...)
}

func (this *extend) GetInt64(key string, def ...int64) int64 {
	return this.GetVar(key).Int64(def...)
}

func (this *extend) GetInt64s(key string, def ...[]int64) []int64 {
	return this.GetVar(key).Int64s(def...)
}

func (this *extend) GetUint(key string, def ...uint) uint {
	return this.GetVar(key).Uint(def...)
}

func (this *extend) GetUint8(key string, def ...uint8) uint8 {
	return this.GetVar(key).Uint8(def...)
}

func (this *extend) GetUint16(key string, def ...uint16) uint16 {
	return this.GetVar(key).Uint16(def...)
}

func (this *extend) GetUint32(key string, def ...uint32) uint32 {
	return this.GetVar(key).Uint32(def...)
}

func (this *extend) GetUint64(key string, def ...uint64) uint64 {
	return this.GetVar(key).Uint64(def...)
}

func (this *extend) GetFloat32(key string, def ...float32) float32 {
	return this.GetVar(key).Float32(def...)
}

func (this *extend) GetFloat64(key string, def ...float64) float64 {
	return this.GetVar(key).Float64(def...)
}

func (this *extend) GetBool(key string, def ...bool) bool {
	return this.GetVar(key).Bool(def...)
}

func (this *extend) GetDuration(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Duration(def...)
}

func (this *extend) GetMicrosecond(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Microsecond(def...)
}

func (this *extend) GetMillisecond(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Millisecond(def...)
}

func (this *extend) GetSecond(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Second(def...)
}

func (this *extend) GetMinute(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Minute(def...)
}

func (this *extend) GetHour(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Hour(def...)
}

func (this *extend) GetErr(key string, def ...error) error {
	return this.GetVar(key).Err(def...)
}

func (this *extend) GetGMap(key string, def ...map[string]interface{}) map[string]interface{} {
	return this.GetVar(key).GMap(def...)
}

func (this *extend) GetSMap(key string, def ...map[string]string) map[string]string {
	return this.GetVar(key).SMap(def...)
}

func (this *extend) GetIMap(key string, def ...map[interface{}]interface{}) map[interface{}]interface{} {
	return this.GetVar(key).IMap(def...)
}

func (this *extend) GetDMap(key string, def ...interface{}) *Map {
	return this.GetVar(key).DMap(def...)
}
