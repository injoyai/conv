package conv

import "time"

type IGetVar interface {
	GetVar(string) *Var
}

func NewVarExtend(i IGetVar) *VarExtend {
	return &VarExtend{IGetVar: i}
}

type VarExtend struct {
	IGetVar
}

func (this *VarExtend) IsDefault(key string) bool {
	return IsDefault(this.GetVar(key).Value)
}

func (this *VarExtend) IsNumber(key string) bool {
	return IsNumber(this.GetVar(key).Value)
}

func (this *VarExtend) IsString(key string) bool {
	return IsString(this.GetVar(key).Value)
}

func (this *VarExtend) IsInt(key string) bool {
	return IsInt(this.GetVar(key).Value)
}

func (this *VarExtend) IsFloat(key string) bool {
	return IsFloat(this.GetVar(key).Value)
}

func (this *VarExtend) IsBool(key string) bool {
	return IsBool(this.GetVar(key).Value)
}

func (this *VarExtend) IsTime(key string) bool {
	return IsTime(this.GetVar(key).Value)
}

func (this *VarExtend) IsArray(key string) bool {
	return IsArray(this.GetVar(key).Value)
}

func (this *VarExtend) IsPointer(key string) bool {
	return IsPointer(this.GetVar(key).Value)
}

func (this *VarExtend) IsNil(key string) bool {
	return IsNil(this.GetVar(key).Value)
}

func (this *VarExtend) Get(key string) *Var {
	return this.GetVar(key)
}

func (this *VarExtend) GetInterface(key string, def ...interface{}) interface{} {
	return this.GetVar(key).Interface(def...)
}

func (this *VarExtend) GetInterfaces(key string, def ...[]interface{}) []interface{} {
	return this.GetVar(key).Interfaces(def...)
}

func (this *VarExtend) GetByte(key string, def ...byte) byte {
	return this.GetVar(key).Byte(def...)
}

func (this *VarExtend) GetBytes(key string, def ...[]byte) []byte {
	return this.GetVar(key).Bytes(def...)
}

func (this *VarExtend) GetRune(key string, def ...rune) rune {
	return this.GetVar(key).Rune(def...)
}

func (this *VarExtend) GetRunes(key string, def ...[]rune) []rune {
	return this.GetVar(key).Runes(def...)
}

func (this *VarExtend) GetString(key string, def ...string) string {
	return this.GetVar(key).String(def...)
}

func (this *VarExtend) GetStrings(key string, def ...[]string) []string {
	return this.GetVar(key).Strings(def...)
}

func (this *VarExtend) GetInt(key string, def ...int) int {
	return this.GetVar(key).Int(def...)
}

func (this *VarExtend) GetInts(key string, def ...[]int) []int {
	return this.GetVar(key).Ints(def...)
}

func (this *VarExtend) GetInt8(key string, def ...int8) int8 {
	return this.GetVar(key).Int8(def...)
}

func (this *VarExtend) GetInt16(key string, def ...int16) int16 {
	return this.GetVar(key).Int16(def...)
}

func (this *VarExtend) GetInt32(key string, def ...int32) int32 {
	return this.GetVar(key).Int32(def...)
}

func (this *VarExtend) GetInt64(key string, def ...int64) int64 {
	return this.GetVar(key).Int64(def...)
}

func (this *VarExtend) GetInt64s(key string, def ...[]int64) []int64 {
	return this.GetVar(key).Int64s(def...)
}

func (this *VarExtend) Uint(key string, def ...uint) uint {
	return this.GetVar(key).Uint(def...)
}

func (this *VarExtend) Uint8(key string, def ...uint8) uint8 {
	return this.GetVar(key).Uint8(def...)
}

func (this *VarExtend) Uint16(key string, def ...uint16) uint16 {
	return this.GetVar(key).Uint16(def...)
}

func (this *VarExtend) Uint32(key string, def ...uint32) uint32 {
	return this.GetVar(key).Uint32(def...)
}

func (this *VarExtend) Uint64(key string, def ...uint64) uint64 {
	return this.GetVar(key).Uint64(def...)
}

func (this *VarExtend) GetFloat32(key string, def ...float32) float32 {
	return this.GetVar(key).Float32(def...)
}

func (this *VarExtend) GetFloat64(key string, def ...float64) float64 {
	return this.GetVar(key).Float64(def...)
}

func (this *VarExtend) GetBool(key string, def ...bool) bool {
	return this.GetVar(key).Bool(def...)
}

func (this *VarExtend) JsonUnmarshal(key string, ptr interface{}) error {
	return this.GetVar(key).JsonUnmarshal(ptr)
}

func (this *VarExtend) GetDuration(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Duration(def...)
}

func (this *VarExtend) GetMicrosecond(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Microsecond(def...)
}

func (this *VarExtend) GetMillisecond(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Millisecond(def...)
}

func (this *VarExtend) GetSecond(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Second(def...)
}

func (this *VarExtend) GetMinute(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Minute(def...)
}

func (this *VarExtend) GetHour(key string, def ...time.Duration) time.Duration {
	return this.GetVar(key).Hour(def...)
}
