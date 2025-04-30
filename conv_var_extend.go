package conv

import "time"

type ExtendString = Extend[string]

type Extend[T any] interface {
	Getter[T]
	IsDefault(key T) bool
	IsNumber(key T) bool
	IsString(key T) bool
	IsInt(key T) bool
	IsFloat(key T) bool
	IsBool(key T) bool
	IsTime(key T) bool
	IsArray(key T) bool
	IsPointer(key T) bool
	IsNil(key T) bool
	Get(key T) *Var
	GetInterface(key T, def ...any) any
	GetInterfaces(key T, def ...[]any) []any
	GetByte(key T, def ...byte) byte
	GetBytes(key T, def ...[]byte) []byte
	GetRune(key T, def ...rune) rune
	GetRunes(key T, def ...[]rune) []rune
	GetString(key T, def ...string) string
	GetStrings(key T, def ...[]string) []string
	GetInt(key T, def ...int) int
	GetInts(key T, def ...[]int) []int
	GetInt8(key T, def ...int8) int8
	GetInt16(key T, def ...int16) int16
	GetInt32(key T, def ...int32) int32
	GetInt64(key T, def ...int64) int64
	GetInt64s(key T, def ...[]int64) []int64
	GetUint(key T, def ...uint) uint
	GetUint8(key T, def ...uint8) uint8
	GetUint16(key T, def ...uint16) uint16
	GetUint32(key T, def ...uint32) uint32
	GetUint64(key T, def ...uint64) uint64
	GetFloat32(key T, def ...float32) float32
	GetFloat64(key T, def ...float64) float64
	GetBool(key T, def ...bool) bool
	GetDuration(key T, def ...time.Duration) time.Duration
	GetMicrosecond(key T, def ...time.Duration) time.Duration
	GetMillisecond(key T, def ...time.Duration) time.Duration
	GetSecond(key T, def ...time.Duration) time.Duration
	GetMinute(key T, def ...time.Duration) time.Duration
	GetHour(key T, def ...time.Duration) time.Duration
	GetErr(key T, def ...error) error
	GetGMap(key T, def ...map[string]any) map[string]any
	GetSMap(key T, def ...map[string]string) map[string]string
	GetIMap(key T, def ...map[any]any) map[any]any
	GetDMap(key T, def ...any) *Map
}

type (
	StringGetter = Getter[string]
	IntGetter    = Getter[int]
)

type Getter[T any] interface {
	GetVar(T) *Var
}

func NewExtendDefault(i Getter[string]) Extend[string] {
	return NewExtendString(i)
}

func NewExtendString(i Getter[string]) Extend[string] {
	return NewExtend[string](i)
}

func NewExtend[T comparable](i Getter[T]) Extend[T] {
	return &extend[T]{Getter: i}
}

type extend[T any] struct {
	Getter[T]
}

func (this *extend[T]) IsDefault(key T) bool {
	return IsDefault(this.GetVar(key).Val())
}

func (this *extend[T]) IsNumber(key T) bool {
	return IsNumber(this.GetVar(key).Val())
}

func (this *extend[T]) IsString(key T) bool {
	return IsString(this.GetVar(key).Val())
}

func (this *extend[T]) IsInt(key T) bool {
	return IsInt(this.GetVar(key).Val())
}

func (this *extend[T]) IsFloat(key T) bool {
	return IsFloat(this.GetVar(key).Val())
}

func (this *extend[T]) IsBool(key T) bool {
	return IsBool(this.GetVar(key).Val())
}

func (this *extend[T]) IsTime(key T) bool {
	return IsTime(this.GetVar(key).Val())
}

func (this *extend[T]) IsArray(key T) bool {
	return IsArray(this.GetVar(key).Val())
}

func (this *extend[T]) IsPointer(key T) bool {
	return IsPointer(this.GetVar(key).Val())
}

func (this *extend[T]) IsNil(key T) bool {
	return IsNil(this.GetVar(key).Val())
}

func (this *extend[T]) Get(key T) *Var {
	return this.GetVar(key)
}

func (this *extend[T]) GetInterface(key T, def ...any) any {
	return this.GetVar(key).Interface(def...)
}

func (this *extend[T]) GetInterfaces(key T, def ...[]any) []any {
	return this.GetVar(key).Interfaces(def...)
}

func (this *extend[T]) GetByte(key T, def ...byte) byte {
	return this.GetVar(key).Byte(def...)
}

func (this *extend[T]) GetBytes(key T, def ...[]byte) []byte {
	return this.GetVar(key).Bytes(def...)
}

func (this *extend[T]) GetRune(key T, def ...rune) rune {
	return this.GetVar(key).Rune(def...)
}

func (this *extend[T]) GetRunes(key T, def ...[]rune) []rune {
	return this.GetVar(key).Runes(def...)
}

func (this *extend[T]) GetString(key T, def ...string) string {
	return this.GetVar(key).String(def...)
}

func (this *extend[T]) GetStrings(key T, def ...[]string) []string {
	return this.GetVar(key).Strings(def...)
}

func (this *extend[T]) GetInt(key T, def ...int) int {
	return this.GetVar(key).Int(def...)
}

func (this *extend[T]) GetInts(key T, def ...[]int) []int {
	return this.GetVar(key).Ints(def...)
}

func (this *extend[T]) GetInt8(key T, def ...int8) int8 {
	return this.GetVar(key).Int8(def...)
}

func (this *extend[T]) GetInt16(key T, def ...int16) int16 {
	return this.GetVar(key).Int16(def...)
}

func (this *extend[T]) GetInt32(key T, def ...int32) int32 {
	return this.GetVar(key).Int32(def...)
}

func (this *extend[T]) GetInt64(key T, def ...int64) int64 {
	return this.GetVar(key).Int64(def...)
}

func (this *extend[T]) GetInt64s(key T, def ...[]int64) []int64 {
	return this.GetVar(key).Int64s(def...)
}

func (this *extend[T]) GetUint(key T, def ...uint) uint {
	return this.GetVar(key).Uint(def...)
}

func (this *extend[T]) GetUint8(key T, def ...uint8) uint8 {
	return this.GetVar(key).Uint8(def...)
}

func (this *extend[T]) GetUint16(key T, def ...uint16) uint16 {
	return this.GetVar(key).Uint16(def...)
}

func (this *extend[T]) GetUint32(key T, def ...uint32) uint32 {
	return this.GetVar(key).Uint32(def...)
}

func (this *extend[T]) GetUint64(key T, def ...uint64) uint64 {
	return this.GetVar(key).Uint64(def...)
}

func (this *extend[T]) GetFloat32(key T, def ...float32) float32 {
	return this.GetVar(key).Float32(def...)
}

func (this *extend[T]) GetFloat64(key T, def ...float64) float64 {
	return this.GetVar(key).Float64(def...)
}

func (this *extend[T]) GetBool(key T, def ...bool) bool {
	return this.GetVar(key).Bool(def...)
}

func (this *extend[T]) GetDuration(key T, def ...time.Duration) time.Duration {
	return this.GetVar(key).Duration(def...)
}

func (this *extend[T]) GetMicrosecond(key T, def ...time.Duration) time.Duration {
	return this.GetVar(key).Microsecond(def...)
}

func (this *extend[T]) GetMillisecond(key T, def ...time.Duration) time.Duration {
	return this.GetVar(key).Millisecond(def...)
}

func (this *extend[T]) GetSecond(key T, def ...time.Duration) time.Duration {
	return this.GetVar(key).Second(def...)
}

func (this *extend[T]) GetMinute(key T, def ...time.Duration) time.Duration {
	return this.GetVar(key).Minute(def...)
}

func (this *extend[T]) GetHour(key T, def ...time.Duration) time.Duration {
	return this.GetVar(key).Hour(def...)
}

func (this *extend[T]) GetErr(key T, def ...error) error {
	return this.GetVar(key).Err(def...)
}

func (this *extend[T]) GetGMap(key T, def ...map[string]any) map[string]any {
	return this.GetVar(key).GMap(def...)
}

func (this *extend[T]) GetSMap(key T, def ...map[string]string) map[string]string {
	return this.GetVar(key).SMap(def...)
}

func (this *extend[T]) GetIMap(key T, def ...map[any]any) map[any]any {
	return this.GetVar(key).IMap(def...)
}

func (this *extend[T]) GetAMap(key T, def ...map[any]any) map[any]any {
	return this.GetVar(key).AMap(def...)
}

func (this *extend[T]) GetDMap(key T, def ...any) *Map {
	return this.GetVar(key).DMap(def...)
}
