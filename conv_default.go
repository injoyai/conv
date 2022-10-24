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
