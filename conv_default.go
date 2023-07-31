package conv

import "time"

func GetDefault(param ...interface{}) *Var {
	return New(Nil().Interface(param...))
}

func GetDefaultInterface(def interface{}, param ...interface{}) interface{} {
	return GetDefault(param...).Interface(def)
}

func GetDefaultBytes(def []byte, param ...[]byte) []byte {
	return GetDefault(Interfaces(param)...).Bytes(def)
}

func GetDefaultString(def string, param ...string) string {
	return GetDefault(Interfaces(param)...).String(def)
}

func GetDefaultUint8(def uint8, param ...uint8) uint8 {
	return GetDefault(Interfaces(param)...).Uint8(def)
}

func GetDefaultUint(def uint, param ...uint) uint {
	return GetDefault(Interfaces(param)...).Uint(def)
}

func GetDefaultInt(def int, param ...int) int {
	return GetDefault(Interfaces(param)...).Int(def)
}

func GetDefaultInt64(def int64, param ...int64) int64 {
	return GetDefault(Interfaces(param)...).Int64(def)
}

func GetDefaultFloat64(def float64, param ...float64) float64 {
	return GetDefault(Interfaces(param)...).Float64(def)
}

func GetDefaultBool(def bool, param ...bool) bool {
	return GetDefault(Interfaces(param)...).Bool(def)
}

func GetDefaultErr(def error, param ...error) error {
	return GetDefault(Interfaces(param)...).Err(def)
}

func GetDefaultDuration(def time.Duration, param ...time.Duration) time.Duration {
	return GetDefault(Interfaces(param)...).Duration(int(def))
}

/*



 */

func Default(param ...interface{}) *Var {
	return New(Nil().Interface(param...))
}

func DefaultInterface(def interface{}, param ...interface{}) interface{} {
	return Default(param...).Interface(def)
}

func DefaultBytes(def []byte, param ...[]byte) []byte {
	return Default(Interfaces(param)...).Bytes(def)
}

func DefaultString(def string, param ...string) string {
	return Default(Interfaces(param)...).String(def)
}

func DefaultUint8(def uint8, param ...uint8) uint8 {
	return Default(Interfaces(param)...).Uint8(def)
}

func DefaultUint16(def uint16, param ...uint16) uint16 {
	return Default(Interfaces(param)...).Uint16(def)
}

func DefaultUint32(def uint32, param ...uint32) uint32 {
	return Default(Interfaces(param)...).Uint32(def)
}

func DefaultUint64(def uint64, param ...uint64) uint64 {
	return Default(Interfaces(param)...).Uint64(def)
}

func DefaultUint(def uint, param ...uint) uint {
	return Default(Interfaces(param)...).Uint(def)
}

func DefaultInt8(def int8, param ...int8) int8 {
	return Default(Interfaces(param)...).Int8(def)
}

func DefaultInt16(def int16, param ...int16) int16 {
	return Default(Interfaces(param)...).Int16(def)
}

func DefaultInt32(def int32, param ...int32) int32 {
	return Default(Interfaces(param)...).Int32(def)
}

func DefaultInt64(def int64, param ...int64) int64 {
	return Default(Interfaces(param)...).Int64(def)
}

func DefaultInt(def int, param ...int) int {
	return Default(Interfaces(param)...).Int(def)
}

func DefaultFloat32(def float32, param ...float32) float32 {
	return Default(Interfaces(param)...).Float32(def)
}

func DefaultFloat64(def float64, param ...float64) float64 {
	return Default(Interfaces(param)...).Float64(def)
}

func DefaultBool(def bool, param ...bool) bool {
	return Default(Interfaces(param)...).Bool(def)
}

func DefaultErr(def error, param ...error) error {
	return Default(Interfaces(param)...).Err(def)
}

func DefaultDuration(def time.Duration, param ...time.Duration) time.Duration {
	return Default(Interfaces(param)...).Duration(int(def))
}
